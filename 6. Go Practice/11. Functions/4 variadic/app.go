package main

import "fmt"

// Variadic function to calculate the sum of numbers
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// 1. Passing individual arguments
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	// 2. Passing no arguments
	fmt.Println("Sum of no numbers:", sum())

	// 3. Using the spread operator with a slice
	numSlice := []int{4, 5, 6}
	fmt.Println("Sum of 4, 5, 6:", sum(numSlice...))

	// 4. Mixing individual arguments and a spread operator
	numSlice2 := []int{7, 8}
	fmt.Println("Sum of 1, 2, 7, 8:", sum(1, 2, numSlice2...))
}
