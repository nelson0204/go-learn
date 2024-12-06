package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	// can create with key:value, and change order
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	//can create without key:value, but need to be the same order of struct
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("nelson0204@gmail.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUSerName()
	admin.OutputUserDetails()
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUSerName()
	appUser.OutputUserDetails()
}

func getUserData(prompText string) string {
	fmt.Print(prompText)
	var value string
	fmt.Scanln(&value)
	return value
}
