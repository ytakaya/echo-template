package view

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ytakaya/echo-template/controller"
)

type Post struct {
	Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Text string `json:"text"`
}

func GetAllPosts(c echo.Context) error {
	db := controller.OpenMySqlConnection()
	defer db.Close()
	db.AutoMigrate(&Post{})

	var posts []Post
	db.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}
