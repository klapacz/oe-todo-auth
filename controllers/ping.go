package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get access token
// @Tags auth
//
// @Produce json
// @Success 200 {object} loginResponseOK
// @security OAuth2Password
//
// @Router /v1/ping [get]
func Ping(c *gin.Context) {
	value, _ := c.Get("claims")
	c.JSON(http.StatusOK, gin.H{
		"message": value,
	})
}
