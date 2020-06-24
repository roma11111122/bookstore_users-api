package users

import (
	"github.com/roma11111122/bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(user.FirstName)
	user.Email = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
