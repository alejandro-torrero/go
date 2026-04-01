package main

import (
	"fmt"
	"structs_test/user"
)

func main() {
	fmt.Println("Starting...")
	// userFirstName := getUserData("Please enter your first name: ")
	// userLastName := getUserData("Please enter your last name: ")
	// userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user.NewAdmin("correo", "123")

	fmt.Println(appUser.GetFullName())
	// appUser.ClearUserName()
	// fmt.Println(appUser.GetFullName())
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
