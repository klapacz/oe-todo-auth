package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klapacz/oe-todo-auth/models"
	"github.com/klapacz/oe-todo-auth/utils"
	"golang.org/x/crypto/bcrypt"
)

// func HashPassword(password string) string {
// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes)
// }

type loginResponseOK struct {
	AccessToken string `json:"access_token"`
}

var badCredentialsErr = HTTPError{Error: "Bad credentials"}

// @Summary Get access token
// @Tags auth
//
// @Accept multipart/form-data
// @Param username formData string true "Username"
// @Param password formData string true "Password"
//
// @Produce json
// @Success 200 {object} loginResponseOK
// @Failure 401 {object} HTTPError
//
// @Router /v1/auth/access-token [post]
func Login(c *gin.Context) {
	var user models.User

	username, password := c.PostForm("username"), c.PostForm("password")

	// fetch user
	if err := models.DB.Where("email = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, badCredentialsErr)
		return
	}

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, badCredentialsErr)
		return
	}

	// create token
	token, err := utils.CreateToken(user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, loginResponseOK{
		AccessToken: token,
	})
}
