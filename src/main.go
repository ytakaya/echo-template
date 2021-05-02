package main

import (
	"github.com/labstack/echo"

	"github.com/ytakaya/echo-template/view"
)

func main() {
	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func initRouting(e *echo.Echo) {
	e.GET("/posts", view.GetAllPosts)
}
