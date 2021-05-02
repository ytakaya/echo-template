package view

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAllPosts(c echo.Context) error {
	body := map[string]string{
		"message": "hello from go api.",
	}

	return c.JSON(http.StatusOK, body)
}
