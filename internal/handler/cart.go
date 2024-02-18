package handler

import (
	"database/sql"
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	db *sql.DB
}

func NewCartHandler(db *sql.DB) *CartHandler {
	return &CartHandler{db: db}
}

func (h *CartHandler) AddItemToCart(c echo.Context) error {
	var item entity.CartItem
	if err := c.Bind(&item); err != nil {
		return err
	}
	h.db.QueryRow("SELECT * FROM cart_item where product_id = $1", item.ProductId)
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
