package main

import (
	"fmt"
)

// Stage 1: Generate numbers
func generateNumbers(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: Square the numbers
func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Stage 3: Filter out odd numbers
func filterEvenNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func main() {
	// Input numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Build the pipeline
	nums := generateNumbers(numbers...)
	squared := squareNumbers(nums)
	evenNumbers := filterEvenNumbers(squared)

	// Collect and print results
	for result := range evenNumbers {
		fmt.Println(result)
	}
}
