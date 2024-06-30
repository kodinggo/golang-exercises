package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	mid1 := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println("before enter handler (middlware 1)")
			err := next(c)
			log.Println("after enter handler (middlware 1)")
			return err
		}
	}

	mid2 := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println("before enter handler (middlware 2)")
			err := next(c)
			log.Println("after enter handler (middlware 2)")
			return err
		}
	}

	e.GET("/", func(c echo.Context) error {
		log.Println("entering handler")
		return c.String(http.StatusOK, "hi")
	}, mid1, mid2)

	e.Start(":8081")
}
