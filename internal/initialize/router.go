package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	c "go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"
	"net/http"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Before -->AA\n")
		c.Next()
		fmt.Printf("After -->AA\n")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Before -->BB\n")
		c.Next()
		fmt.Printf("After -->BB\n")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC\n")
	c.Next()
	fmt.Println("After --> CC\n")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use middleware
	r.Use(middlewares.AuthenMiddleware())

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
