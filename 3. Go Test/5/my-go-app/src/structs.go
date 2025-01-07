// structs.go
package main

import "fmt"

// Define a struct to represent a person
type Person struct {
    Name    string // Name of the person
    Age     int    // Age of the person
    Address string // Address of the person
}

// Constructor function to create a new Person
func NewPerson(name string, age int, address string) *Person {
    return &Person{Name: name, Age: age, Address: address} // Return a pointer to a new Person
}

// Method to display person's information
func (p *Person) DisplayInfo() {
    fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address) // Print person's info
}

// Define a struct to represent an employee, embedding Person
type Employee struct {
    Person // Embedding Person struct
    JobTitle string // Job title of the employee
}

// Method to display employee's information, including job title
func (e *Employee) DisplayEmployeeInfo() {
    e.DisplayInfo() // Call the DisplayInfo method from the embedded Person struct
    fmt.Printf("Job Title: %s\n", e.JobTitle) // Print job title
}

// Main function to demonstrate structs and methods
func main() {
    // Create a new Person using the constructor
    person := NewPerson("Alice", 30, "123 Main St")
    person.DisplayInfo() // Display person's info

    // Create a new Employee
    employee := Employee{
        Person:  *person, // Embed the person
        JobTitle: "Software Engineer", // Set job title
    }
    employee.DisplayEmployeeInfo() // Display employee's info
}