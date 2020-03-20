package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.azc.ext.hp.com/Argon/argon-auth/server/ldapclient"
	"github.com/ant0ine/go-json-rest/rest"
	"gopkg.in/ldap.v3"
)

var config *Config

// StartServer lauches to handle requests
func StartServer() {
	loadConfig()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	authTokenMiddleware := &AuthTokenMiddleware{
		Realm:         "Token Authentication Realm",
		Timeout:       time.Duration(config.JwtAuthentication.TokenDuration),
		Authenticator: LdapAuthenticate,
	}
	authTokenMiddleware.Init()
	api.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != "/login" && request.URL.Path != "/logout"
		},
		IfTrue: authTokenMiddleware,
	})
	router, err := rest.MakeRouter(
		rest.Post("/login", authTokenMiddleware.LoginHandler),
		rest.Post("/modify-attribute", ModifyAttributes),
		rest.Post("/logout", authTokenMiddleware.LogoutHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	log.Fatal(http.ListenAndServeTLS(":9443", "server.cert", "server.key", api.MakeHandler()))
}

func loadConfig() {
	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Failed to open config file: ", err.Error())
	}

	ok, err := ValidateInput(configJSON, Config{})
	if err != nil || ok == false {
		log.Fatal("Invlid configuration: ", err.Error())
	}

	config = new(Config)
	err = json.Unmarshal(configJSON, config)
	if err != nil {
		log.Fatal("Failed to parse config file: ", err.Error())
	}

}

func getLdapConnection() (*ldap.Conn, error) {
	return ldapclient.Bind(config.Ldap.URL, config.Ldap.BindUser, config.Ldap.BindPassword)
}

// ModifyAttributes performs modification of attributes
func ModifyAttributes(w rest.ResponseWriter, r *rest.Request) {
	// ok, err := ValidateRequestInput(*r.Request, LdapAttr{})
	// if err != nil || ok == false {
	// 	ErrorResponse(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	attrToBeModified := &LdapAttr{}
	//	err = r.DecodeJsonPayload(&attrToBeModified)
	err := ValidateAndDecodeRequestPayload(*r.Request, attrToBeModified)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	conn, err := getLdapConnection()
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "LDAP server is not available")
		return
	}
	defer conn.Close()
	ldapclient.ModifyAttr(conn, attrToBeModified.Dn, attrToBeModified.AttrType, attrToBeModified.AttrVals)
	SuccessResponse(w)
}

// LdapAuthenticate performs ldap authentication
func LdapAuthenticate(username string, password string) bool {
	conn, _ := ldapclient.Bind(config.Ldap.URL, config.Ldap.BindUser, config.Ldap.BindPassword)
	defer conn.Close()
	succ, err := ldapclient.AuthByDN(conn, config.Ldap.BaseDn, username, password)

	if err != nil {
		log.Printf(err.Error())
		return false
	}

	return succ
}
