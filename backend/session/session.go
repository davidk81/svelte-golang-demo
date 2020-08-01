package session

// handles http auth requests for path /session
// jwt token generation based on https://github.com/sohamkamani/jwt-go-example

import (
	"encoding/json"
	"log"
	"time"

	"github.com/davidk81/svelte-golang-demo/backend/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

// cookie key
const sessionToken = "session-token"

// TODO: use a secure key mounted during deployment
var jwtKey = []byte("ja93jalkdf092jlkadfh02h3lkdfiu0293lakndf0923haf93ja1h")

// Credentials struct for demarshalling session post body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims struct for jwt token contents
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// HandleSession entrypoint http request handler
func HandleSession(ctx *fasthttp.RequestCtx) error {
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		return handleMethodPost(ctx)
	case "GET":
		return handleMethodGet(ctx)
	case "DELETE":
		return handleMethodDelete(ctx)
	default:
		ctx.NotFound()
		return nil
	}
}

// logs out user by invalidating session token
func handleMethodDelete(ctx *fasthttp.RequestCtx) error {
	var c fasthttp.Cookie
	c.SetKey(sessionToken)
	c.SetValue("")
	c.SetExpire(time.Now())
	ctx.Response.Header.SetCookie(&c)
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

// authenticates user by checking credentials, and sets session token
// if success, responds with user details in post body
func handleMethodPost(ctx *fasthttp.RequestCtx) error {
	// decode login credentials from body
	var creds Credentials
	err := json.Unmarshal(ctx.Request.Body(), &creds)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil
	}

	// check password
	user := user.Login(creds.Username, creds.Password)
	if user == nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return nil
	}

	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err
	}

	// update cookie
	var c fasthttp.Cookie
	c.SetKey(sessionToken)
	c.SetValue(tokenString)
	c.SetExpire(expirationTime)
	ctx.Response.Header.SetCookie(&c)

	// return user info in response, such as roles
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}

// validates session and returns user info in response body
func handleMethodGet(ctx *fasthttp.RequestCtx) error {
	// TODO: verify token & parse username
	ValidateSession(ctx)
	log.Println(ctx.Request.Header.Cookie(sessionToken))

	// fetch user
	user := user.GetUser("username")
	if user == nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return nil
	}

	// return user info in response, such as roles
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

// ValidateSession and check user has atleast one of the roles
func ValidateSession(ctx *fasthttp.RequestCtx, role ...string) bool {
	// TODO:
	return true
}
