package user

import (
	"errors"
	"time"

	serv "github.com/hanifbg/login_register_v2/service"
	util "github.com/hanifbg/login_register_v2/util/password"
	"github.com/hanifbg/login_register_v2/util/validator"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

type CreateUserData struct {
	Name         string `validate:"required"`
	Email        string `validate:"required"`
	Phone_number string `validate:"required,number"`
	Password     string `validate:"required"`
	Address      string
}

func (s *service) CreateUser(data CreateUserData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	hashedPassword, _ := util.EncryptPassword(data.Password)
	user := NewUser(
		data.Name,
		data.Email,
		data.Phone_number,
		string(hashedPassword),
		data.Address,
		time.Now(),
		time.Now(),
	)

	err = s.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LoginUser(email string, password string) (token string, err error) {
	userData, err := s.repository.LoginUser(email)
	if err != nil {
		return "", err
	}

	if !util.ComparePassword(userData.Password, password) {
		return "", errors.New("wrong credentials")
	}

	return "JWT TOKEN", nil
}
