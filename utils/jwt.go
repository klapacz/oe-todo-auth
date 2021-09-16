package utils

import (
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
