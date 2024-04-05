package protos

import (
	//"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	auth "github.com/johncode/articler/auth"
	db "github.com/johncode/articler/db"
	us "github.com/johncode/articler/users"

	"golang.org/x/net/context"
)

type Server struct{}

// mustEmbedUnimplementedChatServiceServer implements ChatServiceServer.
func (s *Server) mustEmbedUnimplementedArticlerServiceServer() {
	fmt.Println("unimplemented")
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recived message body from client: %s", message.GetBody())
	return &Message{Body: "Hello From the Server"}, nil
}

func (s *Server) Register(ctx context.Context, message *LoginForm) (*Message, error) {

	fullUuid := uuid.New().String()
	userId := fullUuid[:30]
	err := db.CreateUser(userId, message.Username, message.Password)

	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(message.Username)
	if err != nil {
		return nil, err
	}

	return &Message{Body: token}, nil
}

func (s *Server) Login(ctx context.Context, message *LoginForm) (*Message, error) {
	var err error
	var user us.User
	user, err = db.GetUser(message.GetUsername())
	if err != nil {
		panic(err)
	}
	log.Println("mes: ", message.Password)

	passwordInput, err := db.HashPassword(message.Password)
	if err != nil {
		panic(err)
	}
	log.Println("p1", passwordInput)
	log.Println("p2", user.Password)

	if user.Password == passwordInput {
		token, err := auth.CreateToken(user.UserId)
		if err != nil {
			return nil, err
		}

		return &Message{Body: token}, nil
	} else {
		return nil, nil
	}

}

func (s *Server) UpdateUser(ctx context.Context, mes *UpdateUserForm) (*Message, error) {

	claims, err := auth.VerifyToken(mes.Token)
	if err != nil {
		return nil, err
	}

	userId := claims["userId"].(string)

	var userInfo us.User
	userInfo.UserId = userId
	userInfo.Username = mes.Username
	userInfo.Password = mes.Password

	err = db.UpdateUserRepo(userInfo)
	if err != nil {
		return nil, err
	}

	return &Message{Body: "success"}, nil
}

// DeleteUser implements ArticlerServiceServer.
func (s *Server) DeleteUser(ctx context.Context, mes *Message) (*Message, error) {

	claims, err := auth.VerifyToken(mes.Body)
	if err != nil {
		return nil, err
	}

	userId := claims["userId"].(string)

	err = db.DeleteUserRepo(userId)
	if err == nil {
		return nil, err
	}

	return &Message{Body: "success"}, nil
}

// Articles

// CreateArticle implements ArticlerServiceServer.
func (s *Server) CreateArticle(context.Context, *ArticleForm) (*Message, error) {
	panic("unimplemented")
}

// UpdateArticle implements ArticlerServiceServer.
func (s *Server) UpdateArticle(context.Context, *ArticleForm) (*Message, error) {
	panic("unimplemented")
}

// DeleteArticle implements ArticlerServiceServer.
func (s *Server) DeleteArticle(context.Context, *DelateArticleForm) (*Message, error) {
	panic("unimplemented")
}

// GetArticles implements ArticlerServiceServer.
func (s *Server) GetArticles(context.Context, *Message) (*Message, error) {
	panic("unimplemented")
}
