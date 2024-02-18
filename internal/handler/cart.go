package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	db *sql.DB
}

func NewCartHandler(db *sql.DB) *CartHandler {
	return &CartHandler{db: db}
}

func (h *CartHandler) AddItemToCart(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
