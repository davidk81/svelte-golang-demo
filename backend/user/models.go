package user

import (
	"context"
	"errors"
)

// WebUserObject for responding to api request
type WebUserObject struct {
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

// Validate client supplied data
func (user WebUserObject) Validate(ctx context.Context) error {
	if len(user.Name) == 0 {
		return errors.New("Name must not be empty")
	}
	if len(user.Username) < 4 {
		return errors.New("Username must be at least 4 characters")
	}
	if len(user.Password) < 8 {
		return errors.New("Password must be at least 8 characters")
	}
	return nil
}
