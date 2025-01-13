package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
	// Base case: factorial of 0 or 1 is 1
	if n == 0 || n == 1 {
		return 1
	}
	// Recursive case: n * factorial(n-1)
	return n * factorial(n-1)
}

func main() {
	number := 5
	fmt.Printf("The factorial of %d is: %d\n", number, factorial(number))

	// Edge case: factorial of 0
	fmt.Printf("The factorial of 0 is: %d\n", factorial(0))
}
