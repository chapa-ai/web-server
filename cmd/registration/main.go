package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"web-server/cmd/registration/controllers"
	"web-server/pkg/config"
	"web-server/pkg/db"
)

func main() {

	err := db.MigrateDb()
	if err != nil {
		fmt.Errorf("MigrateDb failed: %v", err)
		return
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/register", controllers.SignUp)

	config.GetConfig()

	err = e.Start(config.GetConfig().Port)
	if err != nil {
		return
	}

}
