package user_test

import (
	"errors"
	"os"
	"testing"
	"time"

	serv "github.com/hanifbg/login_register_v2/service"
	"github.com/hanifbg/login_register_v2/service/user"
	userMock "github.com/hanifbg/login_register_v2/service/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id           = 1
	name         = "name"
	email        = "main@main.com"
	phone_number = "082212345678"
	password     = "password"
	address      = "address"
)

var (
	userService user.Service
	userRepo    userMock.Repository

	userData          user.User
	userData2         *user.User
	insertUserData    user.CreateUserData
	invalidInsertData user.CreateUserData
)

func setup() {
	userData = user.NewUser(
		name,
		email,
		phone_number,
		password,
		address,
		time.Now(),
		time.Now(),
	)

	userData2 = &userData

	insertUserData = user.CreateUserData{
		Name:         name,
		Email:        email,
		Phone_number: phone_number,
		Password:     password,
		Address:      address,
	}

	invalidInsertData = user.CreateUserData{
		Name:         name,
		Email:        email,
		Phone_number: "",
		Password:     password,
		Address:      address,
	}

	userService = user.NewService(&userRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	t.Run("Expect insert user success", func(t *testing.T) {
		userRepo.On("CreateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(nil).Once()

		err := userService.CreateUser(insertUserData)
		assert.Nil(t, err)
	})

	t.Run("Expect failed user validattion", func(t *testing.T) {
		userRepo.On("CreateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(serv.ErrInvalidData).Once()
		err := userService.CreateUser(invalidInsertData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
	})

	t.Run("Expect failed user validattion", func(t *testing.T) {
		userRepo.On("CreateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(serv.ErrInternalServerError).Once()
		err := userService.CreateUser(insertUserData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInternalServerError)
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("Expect login user notfound", func(t *testing.T) {
		userRepo.On("LoginUser", mock.AnythingOfType("string")).Return(nil, errors.New("record not found")).Once()

		token, err := userService.LoginUser(email, password)
		assert.Equal(t, token, "")
		assert.Equal(t, err, errors.New("record not found"))
	})

	t.Run("Expect login failed wrong password", func(t *testing.T) {
		userRepo.On("LoginUser", mock.AnythingOfType("string")).Return(&userData, nil).Once()

		token, err := userService.LoginUser(email, password)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New("wrong credentials"))
	})
}
