package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type App struct {
	Name string
}

func (a *App) Run() {
	fmt.Println("Welcome to", a.Name)
	for {
		fmt.Print("Enter a number (or 'exit' to quit): ")
		var input string
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println("Exiting the application.")
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		switch {
		case number < 0:
			fmt.Println("You entered a negative number.")
		case number == 0:
			fmt.Println("You entered zero.")
		default:
			fmt.Println("You entered a positive number.")
		}
	}
}

func main() {
	app := &App{Name: "My Go App"}
	app.Run()
}