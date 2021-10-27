package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
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
		"DB_Username": config.DbUsername,
		"DB_Password": config.DbPassword,
		"DB_Port":     strconv.Itoa(config.DbPort),
		"DB_Host":     config.DbAddress,
		"DB_Name":     config.DbName,
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
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)

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
