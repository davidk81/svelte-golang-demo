package session

// handles http auth requests for path /session
// jwt token generation based on https://github.com/sohamkamani/jwt-go-example

import (
	"encoding/json"
	"errors"
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

// HandleSession entrypoint http request handler for /session
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

// HandleRegister entrypoint http request handler for /register
func HandleRegister(ctx *fasthttp.RequestCtx) error {
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		return handleRegisterUser(ctx)
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
	user, err := user.Login(creds.Username, creds.Password, ctx)
	if err != nil {
		return err
	}
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
	user, err := ValidateSession(ctx)
	if err != nil {
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

// ValidateSession and check user has atleast one of the roles. returns WebUserObject object iff session is valid
func ValidateSession(ctx *fasthttp.RequestCtx, roles ...string) (*user.WebUserObject, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(token.Claims)
	if err != nil {
		return nil, err
	}
	var claims Claims
	err = json.Unmarshal(b, &claims)
	if err != nil {
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("session expired")
	}
	user, err := user.GetWebUserObject(claims.Username, ctx)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return user, nil
	}
	for _, requiredRole := range roles {
		for _, myRole := range user.Roles {
			if requiredRole == myRole {
				return user, nil
			}
		}
	}
	return nil, errors.New("user doesn't have role")
}

func verifyToken(ctx *fasthttp.RequestCtx) (*jwt.Token, error) {
	tokenString := string(ctx.Request.Header.Cookie(sessionToken))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// authenticates user by checking credentials, and sets session token
// if success, responds with user details in post body
func handleRegisterUser(ctx *fasthttp.RequestCtx) error {
	// decode login credentials from body
	var newuser user.WebUserObject
	err := json.Unmarshal(ctx.Request.Body(), &newuser)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil
	}

	// check password
	err = user.Register(&newuser, ctx)
	if err != nil {
		return err
	}

	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: newuser.Username,
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
	newuser.Password = "" // sanitize
	b, err := json.Marshal(newuser)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}
