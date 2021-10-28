package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/hanifbg/login_register_v2/config"
	"github.com/hanifbg/login_register_v2/handler"
	userHandler "github.com/hanifbg/login_register_v2/handler/user"
	userRepo "github.com/hanifbg/login_register_v2/repository/user"
	userService "github.com/hanifbg/login_register_v2/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	configData := config.GetConfig()

	dbConnection := config.NewDatabaseConnection(configData)

	userRepo := userRepo.NewGormDBRepository(dbConnection)

	userService := userService.NewService(userRepo)

	userHandler := userHandler.NewHandler(userService)

	e := echo.New()

	handler.RegisterPath(e, userHandler)
	go func() {
		address := fmt.Sprintf(":%d", configData.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
