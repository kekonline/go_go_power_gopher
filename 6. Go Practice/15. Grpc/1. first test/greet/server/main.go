package main

import (
	"log"
	"net"

	pb "example.com/m/greet/proto"
	"google.golang.org/grpc"
)

var address = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", listener.Addr())

	s := grpc.NewServer()
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
