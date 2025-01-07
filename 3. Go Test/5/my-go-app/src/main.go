package main

import (
	"fmt"
	"log"
	"my-go-app/src" // Importing the custom package where other functionalities are defined
)

func main() {
	// Initialize the application
	fmt.Println("Welcome to My Go App!")

	// Call functions from different files to demonstrate various features of Go
	src.DemonstrateVariables()       // Demonstrates variables, values, and operators
	src.DemonstrateConstants()       // Demonstrates constants and type conversions
	src.HandleUserInput()            // Handles user input
	src.BasicErrorHandling()          // Shows basic error handling
	src.ControlFlowExamples()         // Demonstrates control flow statements
	src.FileOperationsExample()       // Demonstrates file reading and writing
	src.PackageExample()              // Shows package organization and usage
	src.PointerExample()              // Demonstrates passing pointers to functions
	src.StructExample()               // Shows struct usage and methods
	src.InterfaceExample()            // Demonstrates interfaces and type assertions
	src.CollectionsExample()          // Covers arrays, slices, and maps
	src.FunctionsExample()            // Demonstrates functions as values and closures
	src.JSONOperationsExample()       // Handles JSON encoding and decoding
	src.ConcurrencyExample()          // Demonstrates goroutines and channels
	src.DebuggingTestingExample()     // Provides guidance on debugging and testing

	// Indicate that the program has finished executing
	fmt.Println("Program execution completed.")
}