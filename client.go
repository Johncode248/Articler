package main

import (
	"log"

	//chat "github.com/grpc-test2/johncode/protos"
	art "github.com/johncode/articler/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("192.168.136.1:9008", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: ", err)
	}

	defer conn.Close()

	c := art.NewArticlerServiceClient(conn)

	input := art.LoginForm{
		Username: "aleksander123",
		Password: "pass12345",
	}

	response, err := c.Login(context.Background(), &input)
	if err != nil {
		panic(err)
	}
	log.Println(response)

	/*
		input := art.LoginForm{
			Username: "admin",
			Password: "12345",
		}
	*/
	//response, err := c.Login(context.Background(), &input)
	//if err != nil {
	//	panic(err)
	//}

	//log.Println("Response: ", response)

	/*
		message := art.Message{
			Body: "Hello from the client",
		}
		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatal("Error when calling SayHello: %s", err)
		}
		//response2, err := c.SayBey(context.Background(), &response)
	*/
	//log.Println("Response from Server: ", response.Body)

}
