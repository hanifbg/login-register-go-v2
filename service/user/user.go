package user

import "time"

type User struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Email        string `gorm:"unique"`
	Phone_number string `gorm:"unique"`
	Password     string
	Address      string
	Role         int
	Token_hash   string
}

func NewUser(
	name string,
	email string,
	phoneNumber string,
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
		Phone_number: phoneNumber,
		Password:     password,
		Address:      address,
		Role:         1,
		Token_hash:   "",
	}
}
