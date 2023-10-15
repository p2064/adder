package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/p2064/adder/handlers"
	"github.com/p2064/adder/proto"
	"github.com/p2064/pkg/config"
	"github.com/p2064/pkg/logs"
	"google.golang.org/grpc"
)

func main() {
	logs.InfoLogger.Print("Start adder")
	if config.Status != config.GOOD {
		logs.ErrorLogger.Fatal("failed to get config")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("ADDER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAdderServiceServer(grpcServer, &handlers.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
