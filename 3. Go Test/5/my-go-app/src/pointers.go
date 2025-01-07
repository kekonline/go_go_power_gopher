// pointers.go
package main

import "fmt"

// This function demonstrates how to pass a pointer to a function
// and modify the value of the variable it points to.
func modifyValue(val *int) {
    // Dereference the pointer and change the value it points to
    *val = 20
}

// This function demonstrates how to use pointers for data mutation
func main() {
    // Declare an integer variable
    number := 10

    // Print the original value of the variable
    fmt.Println("Original value:", number)

    // Pass the address of the variable to the function
    modifyValue(&number)

    // Print the modified value of the variable
    fmt.Println("Modified value:", number)

    // Demonstrating nil pointers
    var ptr *int // Declare a pointer without initializing it
    if ptr == nil {
        fmt.Println("Pointer is nil")
    } else {
        fmt.Println("Pointer points to:", *ptr)
    }

    // Using a pointer to a struct
    type Person struct {
        Name string
        Age  int
    }

    // Create a new Person instance
    person := Person{Name: "Alice", Age: 30}

    // Create a pointer to the Person instance
    personPtr := &person

    // Modify the struct through the pointer
    personPtr.Age = 31

    // Print the modified struct
    fmt.Println("Updated Person:", person)
}