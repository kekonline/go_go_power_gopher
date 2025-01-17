package main

import (
	"fmt"
	"log"

	example "example.com/m/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Simple example
	simpleMsg := &example.SimpleMessage{
		Name: "Alice",
		Age:  30,
	}

	// Serialize the simple message
	simpleData, err := proto.Marshal(simpleMsg)
	if err != nil {
		log.Fatalf("Failed to marshal simple message: %v", err)
	}

	// Deserialize the simple message
	var newSimpleMsg example.SimpleMessage
	err = proto.Unmarshal(simpleData, &newSimpleMsg)
	if err != nil {
		log.Fatalf("Failed to unmarshal simple message: %v", err)
	}

	fmt.Printf("Simple Message - Name: %s, Age: %d\n", newSimpleMsg.Name, newSimpleMsg.Age)

	// Complex example
	complexMsg := &example.ComplexMessage{
		Title: "Complex Example",
		Items: []*example.SimpleMessage{
			{Name: "Bob", Age: 25},
			{Name: "Carol", Age: 28},
		},
		Details: &example.SimpleMessage{
			Name: "Dave",
			Age:  35,
		},
	}

	// Serialize the complex message
	complexData, err := proto.Marshal(complexMsg)
	if err != nil {
		log.Fatalf("Failed to marshal complex message: %v", err)
	}

	// Deserialize the complex message
	var newComplexMsg example.ComplexMessage
	err = proto.Unmarshal(complexData, &newComplexMsg)
	if err != nil {
		log.Fatalf("Failed to unmarshal complex message: %v", err)
	}

	fmt.Printf("Complex Message - Title: %s\n", newComplexMsg.Title)
	for _, item := range newComplexMsg.Items {
		fmt.Printf("Item - Name: %s, Age: %d\n", item.Name, item.Age)
	}
	fmt.Printf("Details - Name: %s, Age: %d\n", newComplexMsg.Details.Name, newComplexMsg.Details.Age)
}
