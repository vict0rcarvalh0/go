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

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example:com", "password")

	admin.OutputUserData()
	admin.ClearUserName()
	admin.OutputUserData()

	appUser.OutputUserData()
	appUser.ClearUserName()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}