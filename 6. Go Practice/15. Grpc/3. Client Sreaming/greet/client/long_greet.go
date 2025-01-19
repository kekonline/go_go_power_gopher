package main

import (
	"context"
	"log"
	"time"

	pb "example.com/m/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "A"},
		{FirstName: "B"},
		{FirstName: "C"},
		{FirstName: "D"},
		{FirstName: "E"},
		{FirstName: "F"},
		{FirstName: "G"},
		{FirstName: "H"},
		{FirstName: "I"},
		{FirstName: "J"},
		{FirstName: "K"},
		{FirstName: "L"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v", req)
		stream.Send(req)

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("LongGreet: %s", res.Result)
}
