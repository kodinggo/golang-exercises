package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bodyRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type queryString struct {
	Sort    string `query:"sort"`
	Keyword string `query:"keyword"`
}

type bodyResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	e := echo.New()

	g := e.Group("/users")
	g.GET("", findAllUsers)
	g.POST("", createUser)
	g.GET("/:id", findUserByID)

	g = e.Group("/products")
	g.GET("", findAllProducts)
	g.POST("", createProduct)

	e.Start(":7723")
}

func findAllUsers(c echo.Context) error {
	var qs queryString
	if err := c.Bind(&qs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(qs.Keyword, qs.Sort)

	resp := bodyResponse{
		Name:    "John",
		Address: "Jl. A",
	}
	// resp := map[string]any{
	// 	"name":    "John",
	// 	"address": "Jl. A",
	// }

	return c.JSON(http.StatusOK, resp)
}

func createUser(c echo.Context) error {
	var body bodyRequest
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, body.Name+body.Address)
}

func findUserByID(c echo.Context) error {
	id := c.Param("id")
	qSort := c.QueryParam("sort")
	return c.String(http.StatusOK, "find user ID="+id+", sort="+qSort)
}

func findAllProducts(c echo.Context) error {
	return c.String(http.StatusOK, "find all users")
}

func createProduct(c echo.Context) error {
	return c.String(http.StatusOK, "create user")
}
