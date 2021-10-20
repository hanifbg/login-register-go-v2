package request

import "github.com/hanifbg/login_register_v2/service/user"

type CreateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	Address      string `json:"address"`
}

func (req *CreateUserRequest) ConvertToUserData() *user.CreateUserData {
	var data user.CreateUserData

	data.Name = req.Name
	data.Email = req.Email
	data.Phone_number = req.Phone_number
	data.Password = req.Password
	data.Address = req.Address

	return &data
}
