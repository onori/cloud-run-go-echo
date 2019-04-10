package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/:param", func(c echo.Context) error {
		param := c.Param("param")
		return c.String(http.StatusOK, param)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
