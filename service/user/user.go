package user

import "time"

type User struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Email        string
	Phone_number string
	Password     string
	Address      string
	Role         int
	Token_hash   string
}

func NewUser(
	name string,
	email string,
	phone_number string,
	password string,
	address string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:           0,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    nil,
		Name:         name,
		Email:        email,
		Phone_number: phone_number,
		Password:     password,
		Address:      address,
		Role:         1,
		Token_hash:   "",
	}
}
