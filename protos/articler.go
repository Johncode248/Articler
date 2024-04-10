package protos

import (
	//"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	art "github.com/johncode/articler/article"
	auth "github.com/johncode/articler/auth"
	db "github.com/johncode/articler/db"
	us "github.com/johncode/articler/users"

	"golang.org/x/net/context"
)

type Server struct{}

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
		log.Println("3 ", err)

		return nil, err
	}

	token, err := auth.CreateToken(userId)
	if err != nil {
		log.Println("4 ", err)
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

func (s *Server) GetAuthorByID(ctx context.Context, message *Message) (*LoginForm, error) {

	claims, err := auth.VerifyToken(message.Body)
	if err != nil {
		return nil, err
	}

	userId := claims["userId"].(string)

	log.Println(userId, " TO JEST ID")

	name, pass := db.GetAuthorRepo(userId)

	return &LoginForm{Username: name, Password: pass}, nil
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

func (s *Server) CreateArticle(ctx context.Context, input *ArticleForm) (*Message, error) {

	claims, err := auth.VerifyToken(input.Token)
	if err != nil {
		return nil, err
	}

	authorId := claims["userId"].(string)
	//fmt.Println("aut ID:  ", a)

	fullUuid := uuid.New().String()
	articleId := fullUuid[:30]

	var articleInput art.Article
	articleInput.Id = articleId
	articleInput.Title = input.Title
	articleInput.Description = input.Content
	articleInput.Summary = input.ShortContent
	articleInput.TimePublished = time.Now()
	articleInput.Author = authorId

	err = db.CreateArticleRepo(articleInput)
	if err != nil {
		return nil, err
	}

	return &Message{Body: "success"}, nil
}

func (s *Server) UpdateArticle(ctx context.Context, input *ArticleForm) (*Message, error) {

	_, err := auth.VerifyToken(input.Token)
	if err != nil {
		return nil, err
	}

	var articleInfo art.Article
	articleInfo.Id = input.ArticleId
	articleInfo.Title = input.Title
	articleInfo.Summary = input.ShortContent
	articleInfo.Description = input.Content

	err = db.UpdateArticleRepo(&articleInfo)
	if err != nil {
		panic(err)
	}

	return &Message{Body: "success"}, nil
}

func (s *Server) DeleteArticle(ctx context.Context, input *DelateArticleForm) (*Message, error) {

	claims, err := auth.VerifyToken(input.Token)
	if err != nil {
		return nil, err
	}

	authorId := claims["userId"].(string)
	log.Println(authorId)

	err = db.DeleteArticleRepo(input.IdArticle)
	if err != nil {
		return nil, err
	}

	return &Message{Body: "success"}, nil
}

// page WAZNE !!!!!
func (s *Server) GetArticles(ctx context.Context, input *Message) (*ListArticles, error) {
	var err error
	var authorToken string = input.Body

	if input.Body != "all" {
		claims, err := auth.VerifyToken(input.Body)
		if err != nil {
			return nil, err
		}
		authorToken = claims["userId"].(string)
	}

	var articles []art.Article

	articles, err = db.GetArticlesRepo(authorToken)
	if len(articles) == 0 {
		fmt.Println(err)
		return nil, nil
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list := ListArticles{}

	for _, articler := range articles {

		var articlesPart ArticleForm
		articlesPart.ArticleId = articler.Id
		articlesPart.Title = articler.Title
		articlesPart.Content = articler.Description
		articlesPart.ShortContent = articler.Summary
		articlesPart.CreateTime = articler.TimePublished.Local().Format("2006-01-02 15:04:05") //Format("2006-01-02 15:04:05") //articler.TimePublished.String()
		articlesPart.AuthorId, _ = db.GetAuthorRepo(articler.Author)

		list.Arts = append(list.Arts, &articlesPart)
	}

	return &list, nil
}
