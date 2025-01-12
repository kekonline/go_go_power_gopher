package main

import (
	"errors"
	"fmt"
	"time"
)

// als the upper case letter is used the struct is available to other packages also the fields are available to other packages
//define a struct caing defines if it is available to other packages
type User struct {
	ID int
	FirstName string
	LastName string
}

// adding a method to the struct
func (u *User) Greeting()  {
	 fmt.Println("Dear", u.FirstName, u.LastName)
}

// adding mutationmethod to the struct
func (u *User) ClearNmae(n string) {
	u.FirstName = ""
}

// adding methos to create a new user
func NewUser(id int, firstName, lastName string) (*User, error) {
	if id == 0 {
		return nil, errors.New("Invalid user ID")
	}
	return &User{ID: id, FirstName: firstName, LastName: lastName}, nil
}

type Admin struct {
	User // embed the User struct into the Admin struct to make the Admin struct with anonymous embedding
	// User User // embed the User struct into the Admin struct to make the Admin struct without anonymous embedding
	Level string
}

func newAdmin(id int, firstName, lastName, level string) (*Admin, error) {
	if id == 0 {
		return nil, errors.New("Invalid user ID")
	}
	if level == "" {
		return nil, errors.New("Invalid admin level")
	}
	return &Admin{User: User{ID: id, FirstName: firstName, LastName: lastName}, Level: level}, nil
}

func (a *Admin) Greeting() {
	fmt.Println("Dear", a.FirstName, a.LastName, "Level:", a.Level)
}

//defina a nested struct
type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
	createDate time.Time	
}



func main() {
fmt.Println("Structs in Go")

// create a new user
user1 := User{ID: 1, FirstName: "John", LastName: "Doe"}
user2 := User{2, "Jane", "Smith"}

//create a new group with the user
group := Group{
	role: "admin",
	users: []User{user1},
	newestUser: user1,
	spaceAvailable: true,
	createDate: time.Now(),

}

admin1 := Admin{
	User: User{ID: 1, FirstName: "John", LastName: "Doe"},
	Level: "super",
}

group.users = append(group.users, user2)

fmt.Println(group)
fmt.Printf("Group: %+v\n", group)
// +v prints the values of the struct
fmt.Printf("The group users are: %+v\n", group.users)

prinUser(&user1)
user1.Greeting()
user1.ClearNmae("John")
user1.Greeting()
admin1.Greeting()
// call the method from the embedded struct with anonymous embedding
admin1.User.Greeting()

//create a new user using the method

user3, err := NewUser(3, "James", "Bond")
if err != nil {
	fmt.Println("Error creating user:", err)
	return
}
fmt.Printf("New user: %+v\n", user3)

}

//function to print the user with a pointer no need to use *u

func prinUser(u *User) {
	fmt.Printf("ID: %d\nFirstName: %s\nLastName: %s\n", (*u).ID, u.FirstName, u.LastName)
}