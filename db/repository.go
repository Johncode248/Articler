package db

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	us "github.com/johncode/articler/users"
)

func CreateUser(userId string, name string, password string) error {

	ins, err := dbUsers.Prepare("INSERT INTO `db1`.`users_table` (`userID`,`name`, `password`) VALUES(?, ?, ?);")
	if err != nil {
		return err
	}
	defer ins.Close()

	password, err = HashPassword(password)
	if err != nil {
		panic(err)
	}

	res, err := ins.Exec(userId, name, password)

	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		log.Fatal(err)
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
	upStmt := "UPDATE `db1`.`users_table` SET `name` = ?, `password` = ? WHERE (`userId` = ?);"
	_, err := dbUsers.Exec(upStmt, userInfo.Username, userInfo.Password, userInfo.UserId)
	if err != nil {
		return err
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
	/*
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		return string(hashedPassword), nil
	*/

	// Tworzenie instancji haszera SHA256
	hasher := sha256.New()

	// Konwertowanie hasła na bajty i haszowanie
	hasher.Write([]byte(password))

	// Pobieranie zsumowanych danych jako bajtów
	hashedPassword := hasher.Sum(nil)

	// Konwersja zsumowanych danych na szesnastkową reprezentację stringa
	hashedPasswordString := hex.EncodeToString(hashedPassword)

	return hashedPasswordString, nil
}
