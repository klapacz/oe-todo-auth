package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klapacz/oe-todo-auth/controllers"
	_ "github.com/klapacz/oe-todo-auth/docs"
	"github.com/klapacz/oe-todo-auth/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title oe-todo auth API
// @version 0.1.0
// @description Manage authorization

// @contact.name Korneliusz Łapacz
// @contact.url http://klapacz.dev
// @contact.email lapacz.kornel@gmail.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /v1/auth/access-token
func main() {
	models.SetupDatabase()

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/auth/access-token", controllers.Login)
	}

	r.Run()
}
