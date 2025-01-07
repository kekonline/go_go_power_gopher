package models

import (
	"fmt"
)

// User represents a user in the application
type User struct {
	ID    int
	Name  string
	Email string
}

// NewUser is a constructor function for creating a new User
func NewUser(id int, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// UpdateEmail updates the user's email
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

// UserList is a custom type for a slice of User
type UserList []User

// AddUser adds a new user to the UserList
func (ul *UserList) AddUser(user User) {
	*ul = append(*ul, user)
}

// GetUserByID retrieves a user by ID
func (ul UserList) GetUserByID(id int) (*User, error) {
	for _, user := range ul {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// UserInterface defines methods for user operations
type UserInterface interface {
	UpdateEmail(newEmail string)
}

// Type assertion example
func PrintUserEmail(ui UserInterface) {
	if user, ok := ui.(*User); ok {
		fmt.Println("User Email:", user.Email)
	} else {
		fmt.Println("Not a User")
	}
}

// Type switch example
func DescribeUser(ui interface{}) {
	switch u := ui.(type) {
	case *User:
		fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
	default:
		fmt.Println("Unknown type")
	}
}