package main

import (
	"echo-handler/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userHandler := handler.NewUserHandler()
	userHandler.RegisterRoutes(e)
	productHandler := handler.NewProductHandler()
	productHandler.RegisterRoutes(e)

	e.Start(":7723")
}
