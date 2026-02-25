package controller

import (
	"fmt"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 2001, []string{uc.userService.GetInfoUser()})
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	type RegisterUserRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, response.ErrCodePramInvalid, err.Error())
		return
	}
	err := uc.userService.RegisterUser(req.Username, req.Email, req.Password)
	fmt.Println("err: ", err)
	if err != response.ErrCodeSuccess {
		response.ErrorResponse(c, err, "")
		return
	}
	response.SuccessResponse(c, response.ErrCodeSuccess, nil)
}
