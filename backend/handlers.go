// based on https://github.com/sohamkamani/jwt-go-example
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func Session(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	if r.Method == "OPTIONS" {
		w.Header().Add("access-control-allow-credentials", "true")
		w.Header().Add("access-control-allow-headers", "Accept,Authorization,Content-Type,If-None-Match")
		w.Header().Add("access-control-allow-methods", r.Header.Get("Access-Control-Request-Method"))
		w.Header().Add("access-control-allow-origin", r.Header.Get("Origin"))
		w.Header().Add("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
		w.Header().Add("access-control-max-age", "86400")
		w.Header().Add("cache-control", "no-cache")
		w.Header().Add("Connection", "keep-alive")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "DELETE" {
		http.SetCookie(w, &http.Cookie{
			Name:    "session-token",
			Value:   "",
			Expires: time.Now(),
		})
		w.Header().Add("access-control-allow-credentials", "true")
		w.Header().Add("access-control-allow-origin", r.Header.Get("Origin"))
		w.Header().Add("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
		w.Header().Add("cache-control", "no-cache")
		w.Header().Add("Connection", "keep-alive")
		w.WriteHeader(http.StatusOK)
		return
	}

	// decode login credentials from body
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check password
	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// update cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session-token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.Header().Add("access-control-allow-credentials", "true")
	w.Header().Add("access-control-allow-origin", r.Header.Get("Origin"))
	w.Header().Add("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
	w.Header().Add("cache-control", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	w.WriteHeader(http.StatusCreated)

	// return user info in response, such as roles
	b, err := json.Marshal(claims)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(b))
}
