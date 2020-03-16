package server

import (
	"log"
	"net/http"

	"github.azc.ext.hp.com/cloud-client/go-ldap-demo/server/ldapclient"
	"github.com/ant0ine/go-json-rest/rest"
)

const (
	SIGN_KEY      = "sign_key"
	URL           = "ldap://192.168.153.164:389"
	BIND_USERNAME = "Administrator@sh.argon"
	BIND_PASSWORD = "myxiaoenen@20191017"
	BASE_DN       = "dc=sh,dc=argon"
)

// StartServer lauches to handle requests
func StartServer() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	authTokenMiddleware := &AuthTokenMiddleware{
		Realm:         "Token Authentication Realm",
		SecretKey:     SIGN_KEY,
		Timeout:       30,
		Authenticator: LdapAuthenticate,
	}
	api.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != "/login"
		},
		IfTrue: authTokenMiddleware,
	})
	router, err := rest.MakeRouter(
		rest.Post("/login", authTokenMiddleware.LoginHandler),
		rest.Post("/modify", ModifyAttributes),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9080", api.MakeHandler()))
}

// LdapAttr struct
type LdapAttr struct {
	Dn       string
	AttrType string
	AttrVals []string
}

// ModifyAttributes performs modification of attributes
func ModifyAttributes(w rest.ResponseWriter, r *rest.Request) {
	attrToBeModified := &LdapAttr{}
	err := r.DecodeJsonPayload(&attrToBeModified)

	if err != nil {
		rest.Error(w, "The payload for modifying attributes is not correct", http.StatusBadRequest)
		return
	}

	ldapclient.ModifyAttr(URL, BIND_USERNAME, BIND_PASSWORD, attrToBeModified.Dn, attrToBeModified.AttrType, attrToBeModified.AttrVals)
	w.WriteHeader(http.StatusOK)
}

// LdapAuthenticate performs ldap authentication
func LdapAuthenticate(username string, password string) bool {
	err := ldapclient.Auth(URL, BASE_DN, BIND_USERNAME, BIND_PASSWORD, username, password)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
