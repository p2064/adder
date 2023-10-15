package main

import (
	"log"
	"net"

	"github.com/p2064/adder/handlers"
	"github.com/p2064/adder/proto"
	"github.com/p2064/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Start adder")
	if config.Status != config.GOOD {
		log.Print("failed to get config")
	}
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAdderServiceServer(grpcServer, &handlers.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
