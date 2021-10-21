package user

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/hanifbg/login_register_v2/common"
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
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.CreateUser(*createUserReq.ConvertToUserData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (handler *Handler) LoginUser(c echo.Context) error {
	createLoginReq := new(request.LoginUserRequest)

	if err := c.Bind(createLoginReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	token, err := handler.service.LoginUser(createLoginReq.Email, createLoginReq.Password)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(token))
}

func (handler *Handler) AuthUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) //conver to jwt.MapClaims

	userID, ok := claims["id"]
	userRole, ok := claims["role"]

	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	data := map[string]string{
		"id":   fmt.Sprintf("%v", userID),
		"role": fmt.Sprintf("%v", userRole),
	}

	return c.JSON(common.NewSuccessResponse(data))
}
