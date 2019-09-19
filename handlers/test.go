package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	defer c.Request().Body.Close()

	return c.JSON(http.StatusOK, "Hello, this is test page")
}
