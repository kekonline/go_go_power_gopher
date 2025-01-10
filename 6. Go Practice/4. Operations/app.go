package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	choice := 1

	for i := 0; i < 10; i++ {
		fmt.Println(i)

// use continue to skip the current iteration
		if i%2 == 0 {
			continue
		}

		// use break to exit the loop
		if i == 5 {
			break
		}

	}
	for {
	fmt.Println("Choose 1 or 2")

	fmt.Print("Enter your choice: ")

	fmt.Scanln(&choice)

	
	if choice == 1 {
		fmt.Println("You chose 1")
	} else if choice == 2 {
		fmt.Println("You chose 2")
	} else {
// write choice to file
	os.WriteFile("choice.txt", []byte(fmt.Sprintf("%d", choice)), 0644)

	// data , _ := os.ReadFile("choice.txt")

	data , err := os.ReadFile("choice.txt")

	if err == nil {
	 fmt.Println("Error reading file")
panic(err)
	//  return
	}
	choice2 := string(data)
		if true {
			fmt.Println("goobye")
			return
		}
		// not printed because of the return statement above
		fmt.Println("this message will not be printed")
	}
	}


	// for someCondition {
	// 	// do something ...
	//   }

// switch exmple

switch choice {
case 1:
	fmt.Println("You chose 1")
case 2:
	fmt.Println("You chose 2")
default:
	fmt.Println("You chose an unknown option")
}


}


func errorTest() (error){
	// error handling
	_, err := os.ReadFile("non_existent_file.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return errors.New("Error reading file")
	}
}