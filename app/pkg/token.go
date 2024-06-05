package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const ExpiredHours = 1

var Key string = "tokenkey"

func GenerateJWTToken(userId int, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userId
	claims["expired_time"] = time.Now().Add(time.Hour * ExpiredHours).Unix()
	claims["role"] = role
	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
