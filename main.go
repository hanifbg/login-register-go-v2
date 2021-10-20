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
	"github.com/hanifbg/login_register_v2/repository/migration"
	userRepo "github.com/hanifbg/login_register_v2/repository/user"
	userService "github.com/hanifbg/login_register_v2/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Host":     os.Getenv("DB_ADDRESS"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	fmt.Println(connectionString)

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func main() {
	config := config.GetConfig()

	dbConnection := newDatabaseConnection(config)

	userRepo := userRepo.NewGormDBRepository(dbConnection)

	userService := userService.NewService(userRepo)

	userHandler := userHandler.NewHandler(userService)

	e := echo.New()

	handler.RegisterPath(e, userHandler)
	fmt.Println(config.AppPort)
	go func() {
		address := fmt.Sprintf("localhost:%d", 8080)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
