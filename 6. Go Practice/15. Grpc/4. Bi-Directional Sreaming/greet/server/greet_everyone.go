package main

import (
	"fmt"
	"io"
	"log"

	pb "example.com/m/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)

		}

		res := fmt.Sprintf("Hello %v", req.FirstName)

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}

}
