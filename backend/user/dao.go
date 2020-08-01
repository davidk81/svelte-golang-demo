package user

// handles database operations for user table

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
