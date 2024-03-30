package protos

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// mustEmbedUnimplementedChatServiceServer implements ChatServiceServer.
func (s *Server) mustEmbedUnimplementedArticlerServiceServer() {
	//panic("unimplemented")
	fmt.Println("unimplemented")
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recived message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server"}, nil
}
