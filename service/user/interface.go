package user

type Service interface {
	CreateUser(CreateUserData) error
	LoginUser(email string, password string) (token string, err error)
}

type Repository interface {
	CreateUser(user User) error
	LoginUser(email string) (*User, error)
}
