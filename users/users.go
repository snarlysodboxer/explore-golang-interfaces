package users

import "fmt"

type UserHandler interface {
	FirstAndLastName() string
}

type User struct {
	FirstName string
	LastName  string
}

func (user *User) FirstAndLastName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func SayHello(user UserHandler) string {
	return fmt.Sprintf("Hello %s", user.FirstAndLastName())
}
