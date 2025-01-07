// interfaces.go
package main

import (
	"fmt"
)

// Define an interface named Shape with a method Area
type Shape interface {
	Area() float64 // Method to calculate area
}

// Define a struct named Rectangle that implements the Shape interface
type Rectangle struct {
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // Calculate area of rectangle
}

// Define a struct named Circle that implements the Shape interface
type Circle struct {
	Radius float64 // Radius of the circle
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius // Calculate area of circle
}

// Function to print the area of a shape
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area()) // Print the area of the shape
}

// Main function to demonstrate the use of interfaces
func main() {
	// Create instances of Rectangle and Circle
	r := Rectangle{Width: 10, Height: 5} // Create a rectangle
	c := Circle{Radius: 7}                // Create a circle

	// Print areas of the shapes
	printArea(r) // Call printArea with Rectangle
	printArea(c) // Call printArea with Circle

	// Demonstrate type assertion
	var shape Shape = r // Assign Rectangle to Shape interface
	if rect, ok := shape.(Rectangle); ok { // Check if shape is of type Rectangle
		fmt.Println("Shape is a Rectangle with width:", rect.Width)
	}

	// Demonstrate type switch
	switch s := shape.(type) {
	case Rectangle:
		fmt.Println("Shape is a Rectangle")
	case Circle:
		fmt.Println("Shape is a Circle")
	default:
		fmt.Println("Unknown shape")
	}
}