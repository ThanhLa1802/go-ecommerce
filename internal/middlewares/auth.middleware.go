package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeTokenInvalid, "Invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
