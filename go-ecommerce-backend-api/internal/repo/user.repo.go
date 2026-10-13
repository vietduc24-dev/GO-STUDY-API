package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (us *UserRepo) GetInfoUser() string {
// 	return "user info duwdcs teo nha"
// }
//INTERFACE_VERSION

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements [IUserRepository].
func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
