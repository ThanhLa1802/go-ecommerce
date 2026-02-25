package service

import (
	"context"
	"encoding/json"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/segmentio/kafka-go"
)

type IUserService interface {
	GetInfoUser() string
	RegisterUser(username, email, password string) int
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

func (us *userService) RegisterUser(username, email, password string) int {
	body := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}
	bodyRequest, _ := json.Marshal(body)
	message := kafka.Message{
		Key:   []byte("register-user"),
		Value: bodyRequest,
	}
	err := global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}
