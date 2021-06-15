package utils

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Reference https://www.howtographql.com/graphql-go/6-authentication/
var SecretKey = []byte("secret")

// generates a jwt token and assign a username to it's claim and return it
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// create a ma p to store our claims
	claims := token.Claims.(jwt.MapClaims)
	// set token claims
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// parse a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
