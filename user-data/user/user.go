package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	age       int
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// Method
func (u User) OutputUserData() {
	fmt.Println("First Name:", u.FirstName)
	fmt.Println("Last Name:", u.LastName)
	fmt.Println("Birth Date:", u.BirthDate)
	fmt.Println("Age:", u.age)
	fmt.Println("Created At:", u.createdAt)
}

// Mutation Method
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "Admin",
			LastName:  "Admin",
			BirthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// Constructor Function
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, fmt.Errorf("please provide all the required fields")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
