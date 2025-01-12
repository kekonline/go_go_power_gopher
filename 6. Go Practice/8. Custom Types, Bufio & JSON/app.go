package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type str string

type person struct {
	Name string
	Age int
}

type personWithTagForJSON struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (s str) print() {
	fmt.Println(s)
}

func main() {
fmt.Println("Custom types in Go")

var s str 
s = "Hello, World!"
s.print()

// reading long string from console
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter a string")
text, err := reader.ReadString('\n')

if err != nil {
	fmt.Println(err)
}

text = strings.TrimSuffix(text, "\n")
text = strings.TrimSuffix(text, "\r")
fmt.Println("You entered: ", text)

// JSON
p1 := person{
	Name: "John",
	Age: 30,
}

p2 := person{
	Name: "John",
	Age: 30,
}

toJSON1, err := json.Marshal(p1)
toJSON2, err := json.MarshalIndent(p2, "", "    ")

if err != nil {
	fmt.Println(err)
}

os.WriteFile("person.json", toJSON1, 0644)
os.WriteFile("personIndent.json", toJSON2, 0644)


}