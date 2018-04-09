package main

import (
	"fmt"
	"github.com/snarlysodboxer/explore-golang-interfaces/users"
)

type CustomUser struct {
	users.User
}

func (customUser *CustomUser) FirstAndLastName() string {
	firstName := customUser.FirstName
	customUser.FirstName = fmt.Sprintf("%s,", customUser.LastName)
	customUser.LastName = firstName

	return customUser.User.FirstAndLastName()
}

func main() {
	user := &users.User{"david", "amick"}

	fmt.Println(users.SayHello(user))

	customUser := &CustomUser{*user}
	fmt.Println(users.SayHello(customUser))
}
