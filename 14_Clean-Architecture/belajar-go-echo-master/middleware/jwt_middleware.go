package middleware

import (
	"belajar-go-echo/konstan"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(konstan.SECRET_JWT))
}
