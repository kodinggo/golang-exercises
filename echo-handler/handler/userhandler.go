package handler

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

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) RegisterRoutes(e *echo.Echo) {
	// Define routes
	g := e.Group("/users")
	g.GET("", h.FindAllUsers)
	g.POST("", h.CreateUser)
	g.GET("/:id", h.FindUserByID)
}

func (h *userHandler) FindAllUsers(c echo.Context) error {
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

func (h *userHandler) CreateUser(c echo.Context) error {
	var body bodyRequest
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, body.Name+body.Address)
}

func (h *userHandler) FindUserByID(c echo.Context) error {
	id := c.Param("id")
	qSort := c.QueryParam("sort")
	return c.String(http.StatusOK, "find user ID="+id+", sort="+qSort)
}
