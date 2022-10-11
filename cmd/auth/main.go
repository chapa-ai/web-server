package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"web-server/cmd/auth/controllers"
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

	e.POST("/api/auth", controllers.Login)
	e.DELETE("/api/auth/:token", controllers.Logout)

	err = e.Start(config.GetConfig().Port)
	if err != nil {
		return
	}

}
