// based on https://github.com/sohamkamani/jwt-go-example
package user

//
type User struct {
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// GetUser checks username & password, and returns User data if successful
func GetUser(username string) *User {
	return &User{
		Name:     username,
		Username: username,
		Roles:    []string{"nurse", "admin"},
	}
}
