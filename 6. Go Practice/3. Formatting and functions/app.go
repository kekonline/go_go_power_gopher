package main

import "fmt"

func main() {
	value1 := 223435.23232

	text1 := fmt.Sprint("3Number: %.0f more numbers %.2f\n", value1, value1)
	text2 :=  fmt.Sprintf("3Number: %.0f more numbers %.2f\n", value1, value1)
	text3 := fmt.Sprintln("3Number: %.0f more numbers %.2f\n", value1, value1)



	fmt.Println("1Number: ", value1)
	fmt.Printf("2Number: %v\n", value1)
	fmt.Printf("2Number: %v more numbers %v\n", value1, value1)
	fmt.Printf("2Number: %f more numbers %T\n", value1, value1)
	fmt.Printf("2Number: %.0f more numbers %.2f\n", value1, value1)



fmt.Println(text1)
fmt.Println(text2)
fmt.Println(text3)

// multi line
fmt.Println(`
	Hello
	World
	`)

test("Hello")

a , b := calculations(10, 5)
c, d := calculations2(10, 5)

}

func test(text string) {
	fmt.Println(text)
}

func calculations (a int, b int) (int, int) {
	return a+b, a-b
}

func calculations2 (a int, b int) (c int, d int) {
	c = a+b
	d = a-b
	
	return 
}