// jwt token generation based on https://github.com/sohamkamani/jwt-go-example
package session

import (
	"encoding/json"
	"fmt"
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
func HandleSession(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		handleMethodPost(ctx)
	case "DELETE":
		handleMethodDelete(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func handleMethodDelete(ctx *fasthttp.RequestCtx) {
	var c fasthttp.Cookie
	c.SetKey(sessionToken)
	c.SetValue("")
	c.SetExpire(time.Now())
	ctx.Response.Header.SetCookie(&c)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func handleMethodPost(ctx *fasthttp.RequestCtx) {
	// decode login credentials from body
	var creds Credentials
	err := json.Unmarshal(ctx.Request.Body(), &creds)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// check password
	user := user.Login(creds.Username, creds.Password)
	if user == nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
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
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
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
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
}
