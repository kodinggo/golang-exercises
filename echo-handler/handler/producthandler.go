package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
}

func NewProductHandler() *productHandler {
	return &productHandler{}
}

func (h *productHandler) RegisterRoutes(e *echo.Echo) {
	// Define routes
	g := e.Group("/products")
	g.GET("", h.FindAllProducts)
	g.POST("", h.CreateProduct)
}

func (h *productHandler) FindAllProducts(c echo.Context) error {
	return c.String(http.StatusOK, "find all users")
}

func (h *productHandler) CreateProduct(c echo.Context) error {
	return c.String(http.StatusOK, "create user")
}
