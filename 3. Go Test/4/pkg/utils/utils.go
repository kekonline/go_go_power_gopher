package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSON reads a JSON file and unmarshals it into the provided interface.
func ReadJSON(filePath string, result interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// WriteJSON encodes the provided interface as JSON and writes it to a file.
func WriteJSON(filePath string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, jsonData, 0644)
}

// HandleUserInput handles user input and returns the input value.
func HandleUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// DeferExample demonstrates the use of defer for deferred execution.
func DeferExample() {
	defer fmt.Println("Deferred execution")
	fmt.Println("Executing function")
}

// SafeDivision performs division and handles division by zero.
func SafeDivision(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}