package user

import (
	"errors"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) GetFullName() string {
	fullName := u.firstName + " " + u.lastName
	return fullName
}

func (u *User) ClearUserName() {
	u.firstName = ""
}

func New(fistName, lastName, birthDate string) (*User, error) {

	if fistName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Invalid data recieved")
	}

	return &User{
		firstName: fistName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Diego",
			lastName:  "Torrero",
			birthDate: "Hoy",
			createdAt: time.Now(),
		},
	}
}
