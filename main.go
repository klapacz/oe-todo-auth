// oe-todo auth API
//
// Manages authorization.
//
//     Schemes: http, https
//     Version: 0.1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Korneliusz Łapacz<lapacz.kornel@gmail.com> http://klapacz.dev
//     SecurityDefinitions:
//       oauth2:
//           type: oauth2
//           tokenUrl: /access-token
//           flow: password
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klapacz/oe-todo-auth/controllers"
	"github.com/klapacz/oe-todo-auth/models"
)

func main() {
	models.SetupDatabase()

	r := gin.Default()
	r.Static("/docs", "./docs")

	r.GET("/ping", controllers.Ping)
	r.POST("/access-token", controllers.Login)
	r.Run()
}
