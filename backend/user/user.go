// based on https://github.com/sohamkamani/jwt-go-example
package user

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

//
type User struct {
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// HandleUser entrypoint http request handler
func HandleUser(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Request.Header.Method()) {
	case "POST":
		handleMethodPost(ctx)
	case "DELETE":
		handleMethodDelete(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

// Login checks username & password, and returns User data if successful
func Login(username, password string) *User {
	// check password
	return &User{
		Name:     username,
		Username: username,
		Roles:    []string{"nurse", "admin"},
	}
}

func handleMethodDelete(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func handleMethodPost(ctx *fasthttp.RequestCtx) {
	// decode post body
	var user User
	err := json.Unmarshal(ctx.Request.Body(), &user)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// return user info in response
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// generate a hashed-and-salted password from plain-text password. return value can be stored in db
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func hashAndSaltPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// compaire plain-text password against a hashed-and-salted password
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
