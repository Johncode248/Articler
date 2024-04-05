package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUsers    *sql.DB
	dbArticles *sql.DB
)

func OpenDatabaseUsers() error {
	var err error
	dbUsers, err = sql.Open("mysql", "root:asekus12345@tcp(127.0.0.1:3306)/db1")
	if err != nil {
		return err
	}
	// Utworzenie tabeli Users
	_, err = dbUsers.Exec(`
	CREATE TABLE IF NOT EXISTS users_table (
		id INT AUTO_INCREMENT PRIMARY KEY,
		userId VARCHAR(45),
		name VARCHAR(45),
		password VARCHAR(100)
	);`)
	if err != nil {
		panic(err)
	}
	return nil
}
func OpenDatabaseArticles() error {
	var err error

	dbArticles, err = sql.Open("mysql", "root:asekus12345@tcp(127.0.0.1:3306)/db1")
	if err != nil {
		panic(err)
	}
	_, err = dbArticles.Exec(`
	CREATE TABLE IF NOT EXISTS article_table (
		id INT AUTO_INCREMENT PRIMARY KEY,
		articleId VARCHAR(45),
		title VARCHAR(45),
		description VARCHAR(100),
		summary VARCHAR(100),
		timePublished VARCHAR(120),
		author VARCHAR(45)
	);`)
	if err != nil {
		panic(err)
	}

	return nil
}

func CreateDatabases() error {

	var err error
	err = OpenDatabaseArticles()
	if err != nil {
		return err
	}
	err = OpenDatabaseUsers()
	if err != nil {
		return err
	}

	return nil
}
