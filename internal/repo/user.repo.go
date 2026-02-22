package repo

type UserRepo struct{}

type IUserRepo interface {
	GetInfoUser() string
}

type userRepo struct {
}

func (*userRepo) GetInfoUser() string {
	return "Thanh La"
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "Thanh La"
// }
