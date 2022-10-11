package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"web-server/cmd/documents/controllers"
	"web-server/pkg/config"
	"web-server/pkg/db"
)

func main() {
	err := db.MigrateDb()
	if err != nil {
		fmt.Sprintf("MigrateDb failed: %v", err)
		return
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/docs", controllers.CreateDocument)
	e.GET("/api/docs", controllers.GetDocumentsList)
	e.GET("/api/docs/:id", controllers.GetDocument)
	e.DELETE("/api/docs/:id", controllers.DeleteDocument)

	err = e.Start(config.GetConfig().Port)
	if err != nil {
		return
	}

}
