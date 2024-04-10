package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/johncode/articler/db"
	art "github.com/johncode/articler/protos"
)

func main() {
	lis, err := net.Listen("tcp", ":9009")
	if err != nil {
		log.Fatal("Failed to listen on port 9008: %v", err)
	}

	err = db.CreateDatabases()
	if err != nil {
		panic(err)
	}
	s := art.Server{}

	grpcServer := grpc.NewServer()

	art.RegisterArticlerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9008: %v", err)
	}

}
