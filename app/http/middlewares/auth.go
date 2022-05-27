package middlewares

import (
	"go-todo/main/app/models"
	"go-todo/main/app/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		tokenId, err := token.ExtractTokenId(c)

		if err != nil {
			c.String(http.StatusBadGateway, err.Error())
			c.Abort()
			return
		}

		user, err := models.GetUserById(tokenId)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
