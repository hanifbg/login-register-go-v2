package user

import (
	"net/http"

	"github.com/hanifbg/login_register_v2/handler/user/request"
	"github.com/hanifbg/login_register_v2/service/user"
	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) CreateUser(c echo.Context) error {
	createUserReq := new(request.CreateUserRequest)

	if err := c.Bind(createUserReq); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	err := handler.service.CreateUser(*createUserReq.ConvertToUserData())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, 2)
	}

	return c.JSON(http.StatusOK, "ok")
}

func (handler *Handler) LoginUser(c echo.Context) error {
	createLoginReq := new(request.LoginUserRequest)

	if err := c.Bind(createLoginReq); err != nil {
		return c.JSON(http.StatusInternalServerError, "cuk")
	}

	token, err := handler.service.LoginUser(createLoginReq.Email, createLoginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
