package main

import "fmt"

// also add functions to custom types

type functionInt func(int, int) int

// 1. Regular function
func add(a, b int) int {
	return a + b
}

// 2. Function that takes another function as an argument
// func applyOperation(a, b int, operation func(int, int) int) int {
func applyOperation(a, b int, operation functionInt) int {
	return operation(a, b)
}

// 3. Function that returns another function
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

func main() {
	// 1. Assigning a function to a variable
	addFunc := add
	fmt.Println("Using function assigned to variable:", addFunc(3, 4))

	// 2. Passing a function as an argument
	result := applyOperation(5, 10, add)
	fmt.Println("Passing function as argument:", result)

	// 3. Returning a function from a function
	double := multiplier(2)
	fmt.Println("Using a function returned by another function:", double(5))

	// 4. Using closures
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println("Closure example:", increment(), increment(), increment())

	// 5. Storing functions in a data structure
	functions := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	}
	fmt.Println("Storing functions in a map:", functions["add"](10, 5), functions["sub"](10, 5))
}
