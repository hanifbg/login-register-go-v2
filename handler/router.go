package handler

import (
	"github.com/hanifbg/login_register_v2/handler/user"
	"github.com/hanifbg/login_register_v2/middleware"
	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *user.Handler) {

	userV1 := e.Group("v1")
	userV1.POST("/register", userHandler.CreateUser)
	userV1.POST("/login", userHandler.LoginUser)

	cobaV1 := e.Group("v1/auth")
	cobaV1.Use(middleware.JWTMiddleware())
	cobaV1.GET("/", userHandler.AuthUser)
}
