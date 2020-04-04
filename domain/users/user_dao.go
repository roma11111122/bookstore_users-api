package users

import (
	"fmt"
	"github.com/roma11111122/bookstore_users-api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewBNotFoundError(fmt.Sprintf("User %d not fount", user.Id))
	}
	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("User with email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
