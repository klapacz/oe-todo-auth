package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klapacz/oe-todo-be-auth/models"
)

func main() {
	models.SetupDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
