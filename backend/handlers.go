// based on https://github.com/sohamkamani/jwt-go-example
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

// TODO: use a secure key mounted during deployment
var jwtKey = []byte("my_secret_key")

// TODO: use real persistence layer
var users = map[string]string{
	"nurse1": "password",
	"nurse2": "password",
	"admin":  "password",
}

// User
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//
type Claims struct {
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func Session(ctx *fasthttp.RequestCtx) {
	if ctx.Request.Header.IsOptions() {
		ctx.Response.Header.Set("access-control-allow-credentials", "true")
		ctx.Response.Header.Set("access-control-allow-headers", "Accept,Authorization,Content-Type,If-None-Match")
		ctx.Response.Header.Set("access-control-allow-methods", string(ctx.Request.Header.Peek("Access-Control-Request-Method")))
		ctx.Response.Header.Set("access-control-allow-origin", string(ctx.Request.Header.Peek("Origin")))
		ctx.Response.Header.Set("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
		ctx.Response.Header.Set("access-control-max-age", "86400")
		ctx.Response.Header.Set("cache-control", "no-cache")
		ctx.Response.Header.Set("Connection", "keep-alive")
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	if ctx.Request.Header.IsDelete() {
		var c fasthttp.Cookie
		c.SetKey("session-token")
		c.SetValue("")
		c.SetExpire(time.Now())
		ctx.Response.Header.SetCookie(&c)

		ctx.Response.Header.Set("access-control-allow-credentials", "true")
		ctx.Response.Header.Set("access-control-allow-origin", string(ctx.Request.Header.Peek("Origin")))
		ctx.Response.Header.Set("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
		ctx.Response.Header.Set("cache-control", "no-cache")
		ctx.Response.Header.Set("Connection", "keep-alive")
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	// decode login credentials from body
	var creds Credentials
	err := json.Unmarshal(ctx.Request.Body(), &creds)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// check password
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	// create jwt token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Name:     "Betty", // TODO:
		Username: creds.Username,
		Roles:    []string{"nurse"}, // TODO:
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
	c.SetKey("session-token")
	c.SetValue(tokenString)
	c.SetExpire(expirationTime)
	ctx.Response.Header.SetCookie(&c)

	ctx.Response.Header.Set("access-control-allow-credentials", "true")
	ctx.Response.Header.Set("access-control-allow-origin", string(ctx.Request.Header.Peek("Origin")))
	ctx.Response.Header.Set("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
	ctx.Response.Header.Set("cache-control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.SetStatusCode(fasthttp.StatusCreated)

	// return user info in response, such as roles
	b, err := json.Marshal(claims)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))

}
