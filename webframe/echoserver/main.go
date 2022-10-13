package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if err := next(c); err != nil {
				c.Error(err)
			}

			fmt.Println("echo middleware")

			return nil
		}
	})

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":6666"))
}

// Handler
func hello(c echo.Context) error {
	fmt.Println("hello")
	return c.String(http.StatusOK, "Hello, World!")
}
