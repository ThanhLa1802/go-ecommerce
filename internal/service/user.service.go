package service

import (
	"go-ecommerce-backend-api/internal/repo"
)

type IUserService interface {
	GetInfoUser() string
}

type userService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
