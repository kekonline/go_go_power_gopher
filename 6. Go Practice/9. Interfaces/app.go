package main

import "fmt"

// interfaces example
type Animal interface {
	Speak() string
}

// embeded interfaces example
type Animal2 interface {
	Animal
	NotSpeak() string
}

type Dog struct {}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Cow struct {}

func (c *Cow) NotSpeak() string {
	return "Moo!"
}

type Hourse struct {}

func (h *Hourse) Speak() string {
	return "Neigh!"
}

func (h *Hourse) NotSpeak() string {
	return "....!"
}

func main() {

	// interfaces example
	var a Animal
	a = &Dog{}
	fmt.Println(a.Speak())

	a = &Cat{}
	fmt.Println(a.Speak())


	// embeded interfaces example
	var a2 Animal2 = &Hourse{}
	fmt.Println(a2.Speak())
	fmt.Println(a2.NotSpeak())



	// example interface error not being implemented
	// var b Animal
	// b = &Cow{}
	// fmt.Println(b.Speak())

	
	
PrintSomething("Hello World")
PrintSomething(1)
PrintSomething(1.1)


result := add(1, 2)
fmt.Println(result)
result2 := add(1.1, 2.2)
fmt.Println(result2)
result3 := add("Hello", "World")
fmt.Println(result3)
// result4 := add(true, false)
// fmt.Println(result4)
result5 := add(1, 2.2)
fmt.Println(result5)
result6 := add(1.1, 2)
fmt.Println(result6)


}

//interface accepts all types
func PrintSomething(toPrint interface{}) {
	switch v := toPrint.(type) {
	case string:
		fmt.Println("String: ", v)
	case int:
		fmt.Println("Int: ", v)
	default:
		fmt.Println("Unknown")
	}

//other way yo check type
	if _, ok := toPrint.(string); ok {
		fmt.Println("String")
	} else if _, ok := toPrint.(int); ok {
		fmt.Println("Int")
	} else {
		fmt.Println("Unknown")
	}

	// _ is used to ignore the value but if you want to use it then you can use it like below
	// if str, ok := toPrint.(string); ok {
	// 	fmt.Println("String: ", str)
	// } else if i, ok := toPrint.(int); ok {
	// 	fmt.Println("Int: ", i)
	// } else {
	// 	fmt.Println("Unknown")
	// }

	//type assertion
	if str, ok := toPrint.(string); ok {
		fmt.Println("String: ", str)
	} else if i, ok := toPrint.(int); ok {
		fmt.Println("Int: ", i)
	} else {
		fmt.Println("Unknown")
	}



}

// generic function
// func add[T any](a, b T) T {
func 		add[T int | float64 | string](a, b T) T {
	return a + b
}