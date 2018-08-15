package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang-rest/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CheckJwtToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	// check jwt token is exist.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(token, token.Claims, err)

	return token.Claims, err
}

func GenerateJWT(u models.User) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	// set token expr time, username, email and id to token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 12 * 1).Unix(),
		"id":    u.ID,
		"name":  u.Username,
		"email": u.Email,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

func CheckPassword(u models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
