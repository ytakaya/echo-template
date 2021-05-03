package main

import (
	"github.com/labstack/echo"

	"github.com/ytakaya/echo-template/controller"
	"github.com/ytakaya/echo-template/view"
)

func main() {
	e := echo.New()
	initDB()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func initDB() {
	db := controller.OpenMySqlConnection()
	db.AutoMigrate(&view.Post{})
}

func initRouting(e *echo.Echo) {
	e.GET("/posts", view.GetAllPosts)
}
