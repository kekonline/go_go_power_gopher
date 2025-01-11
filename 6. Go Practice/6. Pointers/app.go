package main

import "fmt"

func main() {

	// Declaring a variable
	var a int = 10

	// Declaring a pointer
	var p *int

	// Assigning the address of a to the pointer p
	p = &a

	// Accessing the value of a using the pointer p
	fmt.Println("Value of a:", *p)
	multiplyBy10Pointer(p)
	fmt.Println("Value of a after multiplyBy10Pointer:", *p)
	fmt.Println("Value of a after multiplyBy10Pointer:", a)
	multiplyBy10Pointer(&a)
	fmt.Println("Value of a after multiplyBy10Pointer:", *p)
	fmt.Println("Value of a after multiplyBy10Pointer:", a)
}


func multiplyBy10Pointer( number *int) {

*number = *number * 10

}

