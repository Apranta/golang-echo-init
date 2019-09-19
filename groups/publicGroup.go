package groups

import (
	"echostarter_example/handlers"

	"github.com/labstack/echo"
)

func PublicGroup(e *echo.Echo) {
	g := e.Group("")

	// config info
	g.GET("/test", handlers.Test)
}
