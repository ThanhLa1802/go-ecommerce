package repo

import (
	"context"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/database"
)

type UserRepo struct{}

type IUserRepo interface {
	GetInfoUser() string
}

type userRepo struct {
	sqlc *database.Queries
}

func (*userRepo) GetInfoUser() string {
	return "Thanh La"
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (string, error) {
	user, err := ur.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	return user.UsrEmail, nil
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: database.New(global.Mdbc),
	}
}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "Thanh La"
// }
