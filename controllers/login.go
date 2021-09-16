package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/klapacz/oe-todo-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// func HashPassword(password string) string {
// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes)
// }

type loginResponseOK struct {
	AccessToken string `json:"access_token"`
}

const authErrMsg = "Username or password is incorrect"

// @Summary Get access token
// @Tags auth
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} loginResponseOK
// TODO: add error responses
// @Router /v1/auth/access-token [post]
func Login(c *gin.Context) {
	var user models.User

	authErr := func() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": authErrMsg,
		})
	}

	username, password := c.PostForm("username"), c.PostForm("password")

	models.DB.First(&user, "email = ?", username)

	// TODO: find correct method of checking if object was fetched
	if user.Email == "" {
		authErr()
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		authErr()
		return
	}

	token, err := createToken(user)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, loginResponseOK{
		AccessToken: token,
	})
}

var hmacSampleSecret []byte

func createToken(user models.User) (string, error) {
	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        user.ID,
		"token_type": "Bearer",
	})

	return token.SignedString(hmacSampleSecret)
}
