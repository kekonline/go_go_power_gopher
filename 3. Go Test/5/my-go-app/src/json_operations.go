package jsonoperations

import (
	"encoding/json" // Importing the encoding/json package for JSON operations
	"fmt"          // Importing the fmt package for formatted I/O
	"io/ioutil"    // Importing the ioutil package for reading files
	"os"           // Importing the os package for file operations
)

// Person struct to represent a person with Name and Age fields
type Person struct {
	Name string `json:"name"` // JSON tag for Name field
	Age  int    `json:"age"`  // JSON tag for Age field
}

// EncodeToFile encodes a Person object to JSON and writes it to a file
func EncodeToFile(person Person, filename string) error {
	// Convert the Person struct to JSON format
	data, err := json.Marshal(person)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %v", err) // Return error if encoding fails
	}

	// Write the JSON data to the specified file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err) // Return error if writing fails
	}

	return nil // Return nil if successful
}

// DecodeFromFile reads JSON data from a file and decodes it into a Person object
func DecodeFromFile(filename string) (Person, error) {
	var person Person // Declare a Person variable to hold the decoded data

	// Read the JSON data from the specified file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return person, fmt.Errorf("error reading from file: %v", err) // Return error if reading fails
	}

	// Decode the JSON data into the Person struct
	err = json.Unmarshal(data, &person)
	if err != nil {
		return person, fmt.Errorf("error decoding JSON: %v", err) // Return error if decoding fails
	}

	return person, nil // Return the decoded Person object and nil if successful
}

// Example usage of the EncodeToFile and DecodeFromFile functions
func Example() {
	// Create a new Person object
	person := Person{Name: "Alice", Age: 30}

	// Encode the Person object to a file
	err := EncodeToFile(person, "person.json")
	if err != nil {
		fmt.Println("Error:", err) // Print error if encoding fails
		return
	}

	// Decode the Person object from the file
	decodedPerson, err := DecodeFromFile("person.json")
	if err != nil {
		fmt.Println("Error:", err) // Print error if decoding fails
		return
	}

	// Print the decoded Person object
	fmt.Printf("Decoded Person: %+v\n", decodedPerson) // Print the decoded data
}