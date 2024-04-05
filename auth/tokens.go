package auth

//implement token functions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secretkey")

// function for admin
/*
func CreateTokenHandler(userId string) string {

	token, err := CreateToken(userId)
	if err != nil {
		return ""
	}

	return token
}*/

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) (bool, jwt.MapClaims) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "no token", http.StatusUnauthorized)
		return false, nil
	}

	claims, err := VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
		return false, nil
	}

	//username := claims["username"].(string)
	//w.Write([]byte("Token verification successful. User: " + username))
	if err == nil {

		return true, claims
	}
	return true, claims
}

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"userId": username,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
