package main

import "fmt"

func main() {
	// 1. Directly calling an anonymous function
	func() {
		fmt.Println("This is an anonymous function called directly!")
	}()

	// 2. Anonymous function with parameters and immediate invocation
	func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}("Rishi")

	// 3. Assigning an anonymous function to a variable
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Addition using anonymous function:", add(3, 4))

	// 4. Using an anonymous function as a closure
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println("Closure example:", increment(), increment(), increment())

	// 5. Passing an anonymous function as an argument
	execute := func(f func(int, int) int, a, b int) int {
		return f(a, b)
	}
	result := execute(func(x, y int) int { return x * y }, 5, 6)
	fmt.Println("Passing anonymous function as argument:", result)

	// 6. Returning an anonymous function
	multiplier := func(factor int) func(int) int {
		return func(value int) int {
			return value * factor
		}
	}
	double := multiplier(2)
	fmt.Println("Using returned anonymous function:", double(10))
}
