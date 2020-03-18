package server

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	//SecretKey string

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

	//	revokedTokens map[string]int64

	manager SecretKeyManager
}

//Init
func (mw *AuthTokenMiddleware) Init() {
	//	mw.revokedTokens = make(map[string]int64)
	mw.manager = new(SecretKeyManagerFactory).GetSecretKeyManagerManager(MemoryManager)
	mw.manager.ManagerInit()

	// go func() {
	// 	for {
	// 		time.Sleep(time.Minute * 60)
	// 		for tokenString, exp := range mw.revokedTokens {
	// 			if exp < time.Now().Unix() {
	// 				delete(mw.revokedTokens, tokenString)
	// 			}
	// 		}
	// 	}
	// }()
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

type JwtTokenClaims struct {
	Userid string `json:"userid"`
	jwt.StandardClaims
}

// MiddlewareFunc makes authTokenMiddleware implement the Middleware interface.
func (mw *AuthTokenMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	if mw.Authenticator == nil {
		log.Fatal("Authenticator is required")
	}

	// if mw.SecretKey == "" {
	// 	log.Fatal("Secret Key is required")
	// }

	if mw.Authorizator == nil {
		mw.Authorizator = func(userId string, request *rest.Request) bool {
			return true
		}
	}

	return func(w rest.ResponseWriter, r *rest.Request) {
		token, err := mw.extractToken(w, r)
		if err != nil || !mw.validToken(token) {
			mw.unauthorized(w)
			return
		}

		providedUserID := token.Claims.(*JwtTokenClaims).Userid

		if !mw.Authorizator(providedUserID, r) {
			mw.forbidden(w)
			return
		}

		r.Env["REMOTE_USER"] = providedUserID

		handler(w, r)
	}
}

func (mw *AuthTokenMiddleware) validToken(token *jwt.Token) bool {
	if token == nil || !token.Valid {
		return false
	}

	// tokenString, err := token.SignedString([]byte(mw.SecretKey))

	// if err != nil {
	// 	return false
	// }
	// _, exist := mw.revokedTokens[tokenString]
	// if exist {
	// 	return false
	// }

	return true
}

func (mw *AuthTokenMiddleware) extractToken(w rest.ResponseWriter, r *rest.Request) (*jwt.Token, error) {
	claims := &JwtTokenClaims{}
	key, err := mw.getSecretKey(r.Request, request.AuthorizationHeaderExtractor)
	if err != nil || key == "" {
		return nil, err
	}

	token, err := request.ParseFromRequestWithClaims(r.Request, request.AuthorizationHeaderExtractor,
		claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	return token, err
}

func (mw *AuthTokenMiddleware) getSecretKey(req *http.Request, extractor request.Extractor) (string, error) {
	tokeString, err := extractor.ExtractToken(req)
	if err != nil {
		return "", err
	}

	parts := strings.Split(tokeString, ".")
	if len(parts) != 3 {
		return "", nil
	}

	claimstring, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	claims := &JwtTokenClaims{}
	err = json.Unmarshal(claimstring, claims)
	if claims.Userid == "" || err != nil {
		return "", err
	}

	return mw.manager.GetSecretKeyByUser(claims.Userid), nil
}

// LoginHandler performs the authentication, then issue an token.
// It will verify the carried token of the request. If the token is valid, then login will fail.
// Otherwise, login will succeed and existing valid tokens for the userid will be invalid due to the secret key is changed
func (mw *AuthTokenMiddleware) LoginHandler(writer rest.ResponseWriter, request *rest.Request) {
	if token, err := mw.extractToken(writer, request); err == nil && mw.validToken(token) {
		mw.badRequest(writer, "Already logged in")
		return
	}

	account := &Account{}
	err := request.DecodeJsonPayload(&account)

	if err != nil || account.Username == "" || account.Password == "" {
		mw.badRequest(writer, "bad request parameters")
		return
	}

	if !mw.Authenticator(account.Username, account.Password) {
		mw.unauthorized(writer)
		return
	}

	claims := JwtTokenClaims{Userid: account.Username}
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * mw.Timeout).Unix()
	claims.StandardClaims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := mw.manager.CreateSecretKeyByUser(claims.Userid, claims.StandardClaims.ExpiresAt)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		mw.unauthorized(writer)
		return
	}

	SuccessResponseWithPayload(writer, AuthToken{Token: tokenString})

	if mw.PostLoginProcess != nil {
		mw.PostLoginProcess(writer, request)
	}
}

func (mw *AuthTokenMiddleware) unauthorized(writer rest.ResponseWriter) {
	writer.Header().Set("WWW-Authenticate", "AuthToken realm="+mw.Realm)
	ErrorResponse(writer, http.StatusUnauthorized, "Not Authorized")
}

func (mw *AuthTokenMiddleware) badRequest(writer rest.ResponseWriter, errormsg string) {
	ErrorResponse(writer, http.StatusBadRequest, errormsg)
}

func (mw *AuthTokenMiddleware) forbidden(writer rest.ResponseWriter) {
	ErrorResponse(writer, http.StatusForbidden, "Forbidden")
}

// LogoutHandler verifies the carried token in the request. Only if the token is valid
// the secret key for the user will be deleted to invalidate the token.
func (mw *AuthTokenMiddleware) LogoutHandler(w rest.ResponseWriter, r *rest.Request) {
	token, err := mw.extractToken(w, r)
	if err != nil || !mw.validToken(token) {
		mw.badRequest(w, "Did not log in")
		return
	}

	// tokenString, _ := token.SignedString([]byte(mw.SecretKey))

	// mw.revokedTokens[tokenString] = token.Claims.(*JwtTokenClaims).IssuedAt
	mw.manager.DeleteSecretKeyByUser(token.Claims.(*JwtTokenClaims).Userid)

	SuccessResponse(w)
	return
}
