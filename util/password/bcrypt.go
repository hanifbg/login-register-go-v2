package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type utilPassword struct{}

func NewUtil() *utilPassword {
	return &utilPassword{}
}

func (u *utilPassword) EncryptPassword(pass string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

func (u *utilPassword) ComparePassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}

	return false
}
