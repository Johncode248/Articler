package db

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	art "github.com/johncode/articler/article"
	us "github.com/johncode/articler/users"
)

func CreateUser(userId string, name string, password string) error {

	ins, err := dbUsers.Prepare("INSERT INTO `db1`.`users_table` (`userID`,`name`, `password`) VALUES(?, ?, ?);")
	if err != nil {
		return err
	}
	defer ins.Close()

	passwordHashed, err := HashPassword(password)
	if err != nil {
		panic(err)
	}

	_, err = ins.Exec(userId, name, passwordHashed)

	if err != nil {
		return err
	}

	log.Println("Created")

	return nil
}

func GetUser(username string) (us.User, error) {
	var user us.User
	var i int
	row := dbUsers.QueryRow("SELECT * FROM db1.users_table WHERE name = ?;", username)

	err := row.Scan(&i, &user.UserId, &user.Username, &user.Password)
	log.Println("id", i, " userid: ", user.UserId, " Username: ", user.Username, " Password: ", user.Password)
	return user, err
}

func UpdateUserRepo(userInfo us.User) error {

	var upStmt string

	fmt.Println("Updating")
	fmt.Println("username: ", userInfo.Username)
	fmt.Println("pass: ", userInfo.Password)

	if userInfo.Username != "" {
		upStmt = "UPDATE `db1`.`users_table` SET `name` = ? WHERE (`userId` = ?);"
		_, err := dbUsers.Exec(upStmt, userInfo.Username, userInfo.UserId)
		if err != nil {
			return err
		}
	}

	if userInfo.Password != "" {
		hashedPass, err := HashPassword(userInfo.Password)
		if err != nil {
			return err
		}
		upStmt = "UPDATE `db1`.`users_table` SET `password` = ? WHERE (`userId` = ?);"
		_, err = dbUsers.Exec(upStmt, hashedPass, userInfo.UserId)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUserRepo(userId string) error {
	_, err := dbUsers.Exec("DELETE FROM `db1`.`users_table` WHERE (`userId` = ?)", userId)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {

	hasher := sha256.New()

	hasher.Write([]byte(password))

	hashedPassword := hasher.Sum(nil)

	hashedPasswordString := hex.EncodeToString(hashedPassword)

	return hashedPasswordString, nil
}

// Articles

func CreateArticleRepo(article art.Article) error {
	ins, err := dbArticles.Prepare("INSERT INTO `db1`.`article_table` (`articleId`,`title`, `description`, `summary`,`timePublished`, `author`) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer ins.Close()

	res, err := ins.Exec(article.Id, article.Title, article.Description, article.Summary, article.TimePublished, article.Author)

	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		log.Fatal(err)
		return err
	}

	log.Println("Created")

	return nil
}

func DeleteArticleRepo(articleId string) error {
	_, err := dbUsers.Exec("DELETE FROM `db1`.`article_table` WHERE (`articleId` = ?)", articleId)
	if err != nil {
		return err
	}

	return nil
}

func GetArticlesRepo(authorId string, pagination string) ([]art.Article, error) {

	var err error
	page := 1
	pageSize := 8

	if pagination != "" {
		page, err = strconv.Atoi(pagination)
		if err != nil {
			fmt.Println(err)
			fmt.Println("konvercja")
		}
	}

	var query string
	if authorId == "all" {
		query = fmt.Sprintf("SELECT articleId, title, description, summary, timePublished, author FROM db1.article_table ORDER BY timePublished DESC LIMIT %d OFFSET %d", pageSize, (page-1)*pageSize)
	} else {
		query = fmt.Sprintf("SELECT articleId, title, description, summary, timePublished, author FROM db1.article_table WHERE author = '%s' ORDER BY timePublished DESC LIMIT %d OFFSET %d", authorId, pageSize, (page-1)*pageSize)

	}

	// Wykonaj zapytanie do bazy danych
	rows, err := dbArticles.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var articles []art.Article

	for rows.Next() {
		var article art.Article

		var timePublishedString string
		err := rows.Scan(&article.Id, &article.Title, &article.Description, &article.Summary, &timePublishedString, &article.Author) // add time.Published
		if err != nil {
			return nil, err
		}
		article.TimePublished, err = time.Parse("2006-01-02 15:04:05", timePublishedString)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	fmt.Println("Repo:  ", articles)

	return articles, nil
}

func GetAuthorRepo(authorId string) (string, string) {

	fmt.Println("id  ", authorId)
	query := "SELECT name, password FROM `db1`.`users_table` WHERE userId = ?;"
	var name string
	var pass string
	err := dbUsers.QueryRow(query, authorId).Scan(&name, &pass)
	if err != nil {
		return "", ""
	}

	fmt.Println(err)

	return name, pass
}

func UpdateArticleRepo(article *art.Article) error {

	var upStmt string

	if article.Title != "" {
		upStmt = "UPDATE `db1`.`article_table` SET `title` = ? WHERE (`articleId` = ?);"
		_, err := dbArticles.Exec(upStmt, article.Title, article.Id)
		if err != nil {
			return err
		}
	}

	if article.Summary != "" {

		upStmt = "UPDATE `db1`.`article_table` SET `summary` = ? WHERE (`articleId` = ?);"
		_, err := dbArticles.Exec(upStmt, article.Summary, article.Id)
		if err != nil {
			return err
		}
	}

	if article.Description != "" {

		upStmt = "UPDATE `db1`.`article_table` SET `description` = ? WHERE (`articleId` = ?);"
		_, err := dbArticles.Exec(upStmt, article.Description, article.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
