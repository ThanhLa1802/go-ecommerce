package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 2001, []string{"Thanh La"})
}
