package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /ping Ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
