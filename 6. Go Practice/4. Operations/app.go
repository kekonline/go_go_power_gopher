package main

import "fmt"

func main() {
	choice := 1

	fmt.Println("Choose 1 or 2")

	fmt.Print("Enter your choice: ")

	fmt.Scanln(&choice)

	

	if choice == 1 {
		fmt.Println("You chose 1")
	} else if choice == 2 {
		fmt.Println("You chose 2")
	} else {
		fmt.Println("You chose something else")
	}


}