package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "example.com/m/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {

	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	reqs := []*pb.GreetRequest{

		{FirstName: "John"},
		{FirstName: "Mark"},
		{FirstName: "Pam"},
		{FirstName: "Todd"},
		{FirstName: "Amy"},
	}

	waitc := make(chan struct{})

	go func() {

		for _, req := range reqs {

			log.Printf("Sending request: %v", req)
			stream.Send(req)

			time.Sleep(1000 * time.Millisecond)

		}

		stream.CloseSend()

	}()

	go func() {

		for {

			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v", res.Result)

		}

		close(waitc)

	}()

	<-waitc

}
