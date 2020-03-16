package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// AuthTokenMiddleware provides a token authentication implementation. On failure, a 401 HTTP response
//is returned. On success, the wrapped middleware is called, and the userId is made available as
// request.Env["REMOTE_USER"].(string)
type AuthTokenMiddleware struct {

	// Realm name to display to the user. Required.
	Realm string

	// The key to sign the token. Required
	SecretKey string

	// Timeout of token
	Timeout time.Duration

	// Callback function that should perform the authentication of the user based on userId and
	// password. Must return true on success, false on failure. Required.
	Authenticator func(username string, password string) bool

	// Callback function that should perform the authorization of the authenticated user. Called
	// only after an authentication success. Must return true on success, false on failure.
	// Optional, default to success.
	Authorizator func(userId string, r *rest.Request) bool

	// Callback function that should perform postprocess of the login event.
	PostLoginProcess func(w rest.ResponseWriter, r *rest.Request)
}

// Account struct
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthToken struct
type AuthToken struct {
	Token string `json:"token"`
}

// MiddlewareFunc makes authTokenMiddleware implement the Middleware interface.
func (mw *AuthTokenMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	if mw.Authenticator == nil {
		log.Fatal("Authenticator is required")
	}

	if mw.SecretKey == "" {
		log.Fatal("Secret Key is required")
	}

	if mw.Authorizator == nil {
		mw.Authorizator = func(userId string, request *rest.Request) bool {
			return true
		}
	}

	return func(writer rest.ResponseWriter, r *rest.Request) {

		claims := make(jwt.MapClaims)
		token, err := request.ParseFromRequestWithClaims(r.Request, request.AuthorizationHeaderExtractor,
			claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(mw.SecretKey), nil
			})

		if err != nil || !token.Valid {
			mw.unauthorized(writer)
			return
		}

		providedUserID := claims["username"]

		r.Env["REMOTE_USER"] = providedUserID

		handler(writer, r)
	}
}

// LoginHandler performs the authentication, then issue an token
func (mw *AuthTokenMiddleware) LoginHandler(writer rest.ResponseWriter, request *rest.Request) {
	account := &Account{}
	err := request.DecodeJsonPayload(&account)

	if err != nil {
		mw.unauthorized(writer)
		return
	}

	if !mw.Authenticator(account.Username, account.Password) {
		mw.unauthorized(writer)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * mw.Timeout).Unix()
	claims["orig_iat"] = time.Now().Unix()
	claims["id"] = account.Username
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(mw.SecretKey))

	if err != nil {
		mw.unauthorized(writer)
		return
	}

	writer.WriteJson(AuthToken{Token: tokenString})

	if mw.PostLoginProcess != nil {
		mw.PostLoginProcess(writer, request)
	}
}

func (mw *AuthTokenMiddleware) unauthorized(writer rest.ResponseWriter) {
	writer.Header().Set("WWW-Authenticate", "AuthToken realm="+mw.Realm)
	rest.Error(writer, "Not Authorized", http.StatusUnauthorized)
}
