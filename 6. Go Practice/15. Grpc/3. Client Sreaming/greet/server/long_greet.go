package main

import (
	"fmt"
	"io"
	"log"

	pb "example.com/m/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving request: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}

}
