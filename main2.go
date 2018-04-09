package main

import (
	"encoding/json"
	"fmt"
	"github.com/snarlysodboxer/explore-golang-interfaces/users2"
	"strings"
)

type CustomUser struct {
	users.User
}

// Reverse order and add ','
func (customUser *CustomUser) SetFirstAndLastName() {
	customUser.FirstAndLastName = fmt.Sprintf("%s, %s", customUser.LastName, customUser.FirstName)
}

// Modify just before write
func (customUser *CustomUser) Write() error {
	customUser.FirstName = strings.ToUpper(customUser.FirstName)
	customUser.LastName = strings.ToUpper(customUser.LastName)

	// Call original
	err := customUser.User.Write()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Normal usage
	user := &users.User{}
	user.FirstName = "david"
	user.LastName = "amick"
	userHandler, err := users.CreateUser(user)
	if err != nil {
		fmt.Errorf("Error with CreateUser: %#v\n", err)
	}
	fmt.Printf("Added user: %s\n", prettyify(userHandler))

	// Customize
	customUser := &CustomUser{users.User{}}
	customUser.FirstName = "clyde"
	customUser.LastName = "McNab"
	customUserReader, err := users.CreateUser(customUser)
	if err != nil {
		fmt.Errorf("Error with CreateUser: %#v\n", err)
	}
	fmt.Printf("Added custom user: %s\n", prettyify(customUserReader))
}

func prettyify(user users.UserHandler) string {
	pretty, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Errorf("%v\n", err)
	}
	return string(pretty)
}
