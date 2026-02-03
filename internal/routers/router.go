package routers

import (
	"github.com/gin-gonic/gin"
	c "go-ecommerce-backend-api/internal/controller"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1/2026")
	{
		v1.GET("/user", c.NewUserController().GetUserById)
		v1.POST("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong....ping...test",
	})
}
