package main

import (
	"context"
	"io"
	"log"

	pb "example.com/m/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "John",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
