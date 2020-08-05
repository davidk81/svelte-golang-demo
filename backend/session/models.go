package session

import "github.com/dgrijalva/jwt-go"

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
