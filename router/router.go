package router

import (
	"echostarter_example/groups"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	groups.PublicGroup(e)

	return e
}
