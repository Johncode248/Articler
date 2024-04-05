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
	lis, err := net.Listen("tcp", ":9008")
	if err != nil {
		log.Fatal("Failed to listen on port 9008: %v", err)
	}
	//
	p1, _ := db.HashPassword("asekus123")
	p2, _ := db.HashPassword("asekus123")
	p3, _ := db.HashPassword("asekus123")
	log.Println("p1 ", p1)
	log.Println("p2 ", p2)
	log.Println("p3 ", p3)

	//

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
