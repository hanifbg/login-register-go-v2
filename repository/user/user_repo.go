package user

import (
	"fmt"
	"time"

	"github.com/hanifbg/login_register_v2/service/user"

	"gorm.io/gorm"
)

type User struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string `json:"name"  validate:"required"`
	Email       string `json:"email" validate:"required,email" gorm:"type:varchar(20),unique"`
	PhoneNumber string `json:"phone_number" validate:"required,number" gorm:"unique"`
	Password    string `json:"password"  validate:"required"`
	Address     string `json:"address"  validate:"required"`
	Role        int
	Token_hash  string
}

type GormRepository struct {
	DB *gorm.DB
}

func newUserTable(user user.User) *User {

	return &User{
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		user.Name,
		user.Email,
		user.Phone_number,
		user.Password,
		user.Address,
		user.Role,
		user.Token_hash,
	}
}
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (col *User) ToUser() user.User {
	var user user.User

	user.ID = col.ID
	user.Name = col.Name
	user.Email = col.Email
	user.Password = col.Password
	user.CreatedAt = col.CreatedAt
	user.UpdatedAt = col.UpdatedAt
	user.DeletedAt = col.DeletedAt
	user.Address = col.Address
	user.Role = col.Role
	user.Token_hash = col.Token_hash
	user.Phone_number = col.PhoneNumber

	return user
}

func (repo *GormRepository) CreateUser(user user.User) error {
	userData := newUserTable(user)

	err := repo.DB.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormRepository) LoginUser(email string) (*user.User, error) {
	var userData User

	err := repo.DB.Where("email = ?", email).First(&userData).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}
