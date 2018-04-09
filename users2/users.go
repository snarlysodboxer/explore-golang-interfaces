package users

import (
	"fmt"
)

type UserHandler interface {
	SetID()
	SetFirstAndLastName()
	Write() error
}

type User struct {
	FirstName        string
	LastName         string
	ID               int
	FirstAndLastName string
}

func (user *User) SetID() {
	// Pretend to generated an ID
	user.ID = 123456
}

func (user *User) SetFirstAndLastName() {
	user.FirstAndLastName = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func (user *User) Write() error {
	// Pretend to write to a DB
	// if err != nil {
	//     return err
	// }
	return nil
}

func CreateUser(user UserHandler) (UserHandler, error) {
	user.SetID()
	user.SetFirstAndLastName()
	err := user.Write()
	if err != nil {
		return user, err
	}
	return user, nil
}
