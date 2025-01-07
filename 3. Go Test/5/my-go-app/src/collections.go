package main

import (
	"fmt"
)

// This function demonstrates the use of arrays and slices in Go.
func arrayAndSliceDemo() {
	// Declare an array of integers with a fixed size of 5
	var arr [5]int
	// Initialize the array with values
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1 // Assign values 1 to 5
	}
	fmt.Println("Array:", arr)

	// Create a slice from the array
	slice := arr[1:4] // Get a slice from index 1 to 3 (not including 4)
	fmt.Println("Slice:", slice)

	// Append a value to the slice
	slice = append(slice, 6) // Append 6 to the slice
	fmt.Println("Updated Slice:", slice)

	// Demonstrate dynamic slices
	dynamicSlice := make([]int, 0) // Create an empty slice
	for i := 1; i <= 5; i++ {
		dynamicSlice = append(dynamicSlice, i) // Append values 1 to 5
	}
	fmt.Println("Dynamic Slice:", dynamicSlice)

	// Demonstrate slicing operations
	slicedPart := dynamicSlice[1:3] // Get a slice from index 1 to 2
	fmt.Println("Sliced Part:", slicedPart)

	// Demonstrate maps and their mutation
	m := make(map[string]int) // Create a map with string keys and int values
	m["one"] = 1              // Add key-value pairs
	m["two"] = 2
	m["three"] = 3
	fmt.Println("Map:", m)

	// Update a value in the map
	m["two"] = 22
	fmt.Println("Updated Map:", m)

	// Delete a key from the map
	delete(m, "one")
	fmt.Println("Map after deletion:", m)
}

// The main function serves as the entry point of the application.
func main() {
	// Call the demo function to showcase arrays, slices, and maps
	arrayAndSliceDemo()
}