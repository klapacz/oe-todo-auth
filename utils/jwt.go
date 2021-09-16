package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/klapacz/oe-todo-auth/models"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        user.ID,
		"token_type": "Bearer",
	})

	return token.SignedString(JWTSecret)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return JWTSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("error while decoding token: %v", err)
	}
}
