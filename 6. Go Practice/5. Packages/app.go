package main

// get packages from: https://pkg.go.dev/github.com/pallinder/go-randomdata
// go get github.com/pallinder/go-randomdata
// import "github.com/pallinder/go-randomdata"

//import you packages need to include full path like in the module
import (
	"fmt"

	"example.com/packages/anotherPackage"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Hello from main!")
	printFromAnotherFile()
	anotherPackage.PrintFromEvenAnotherFile()
	fmt.Println(randomdata.FirstName(randomdata.Male))
}