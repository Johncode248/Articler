package main

import (
	"log"
	"net"

	//chat "github.com/grpc-test2/johncode/protos"

	"google.golang.org/grpc"

	art "github.com/johncode/articler/protos"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("Failed to listen on port 9001: %v", err)
	}

	s := art.Server{}

	grpcServer := grpc.NewServer()

	art.RegisterArticlerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9001: %v", err)
	}

}
