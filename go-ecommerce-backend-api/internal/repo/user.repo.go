package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (us *UserRepo) GetInfoUser() string {
	return "user info duwdcs teo nha"
}
