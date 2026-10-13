package service

import (
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }
//interface VERSION

type IUserService interface {
	Register(email string, purpore string) int
}

type userService struct {
	userRepo repo.IUserRepository
	//...
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements [IUserService].
func (us *userService) Register(email string, purpore string) int {
	//1. check email exist
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExisted
	}
	//2. hash password
	//3. register
	//4. return
	return response.ErrCodeSuccess
}
