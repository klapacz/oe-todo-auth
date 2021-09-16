package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/klapacz/oe-todo-auth/controllers"
	"github.com/klapacz/oe-todo-auth/utils"
)

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.Request.Header["Authorization"][0], "Bearer ")
	claims, err := utils.ParseToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.HTTPError{
			Error: "Invalid token",
		})
		return
	}

	c.Set("claims", claims)
	c.Next()
}
