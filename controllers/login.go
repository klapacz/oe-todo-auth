package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func HashPassword(password string) string {
// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes)
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// swagger:parameters credentails login
type loginRequest struct {
	// example: user@example.com
	// in: formData
	// required: true
	Username string `form:"username" json:"username" binding:"required"`
	// in: formData
	// required: true
	Password string `form:"password" json:"password" binding:"required"`
}

// swagger:response loginResponseOK
type loginResponseOK struct {
	// in: body
	Body struct {
		AccessToken string `json:"accessToken"`
	}
}

// swagger:response GenericError
type genericError struct {
	// in: body
	Body struct {
		Error string `json:"error"`
	}
}

const authErrMsg = "Username or password is incorrect"

// swagger:route POST /access-token auth login
//
//     Responses:
//       200: loginResponseOK
//       422: GenericError
//       401: GenericError
//     Consumes:
//       - application/x-www-form-urlencoded
//
func Login(c *gin.Context) {
	// var user models.User

	// authErr := func() {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": authErrMsg,
	// 	})
	// }

	// if err := c.ShouldBindJSON(&credentails); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// }

	// models.DB.First(&user, "email = ?", credentails.Username)

	// // TODO: find correct method of checking if object was fetched
	// if user.Email == "" {
	// 	authErr()
	// 	return
	// }
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentails.Password))
	// if err != nil {
	// 	authErr()
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"access_token": "hello",
	})
}
