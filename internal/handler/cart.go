package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	db *sql.DB
}

func NewCartHandler(db *sql.DB) *CartHandler {
	return &CartHandler{db: db}
}

func (h *CartHandler) FetchCart(c echo.Context) error {
	userId := c.Get("userId")
	cartItems := []entity.CartItem{}
	totalPrice := 0.0
	rows, err := h.db.Query("SELECT c.id, c.quantity, c.total_price, p.id, p.title, p.image_url FROM cart_items c, products p WHERE c.product_id = p.id AND user_id = $1", userId)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var item entity.CartItem
		if err = rows.Scan(&item.Id, &item.Quantity, &item.TotalPrice, &item.Product.Id, &item.Product.Title, &item.Product.ImageUrl); err != nil {
			return err
		}
		cartItems = append(cartItems, item)
		totalPrice += item.TotalPrice
	}
	return c.JSON(http.StatusOK, map[string]any{"items": cartItems, "totalPrice": totalPrice})
}

func (h *CartHandler) AddItemToCart(c echo.Context) error {
	var req entity.CartItemRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	var item entity.CartItem
	var cartId string
	row := h.db.QueryRow("SELECT quantity, total_price FROM cart_items where product_id = $1", req.ProductId)
	if err := row.Scan(&item.Quantity, &item.TotalPrice); err != nil {
		if err == sql.ErrNoRows {
			// add item to cart
			newCartId := uuid.NewString()
			userId := c.Get("userId").(string)
			if _, err := h.db.Exec("INSERT INTO cart_items (id, user_id, product_id, quantity, total_price) VALUES ($1, $2, $3, $4, $5)", newCartId, userId, req.ProductId, req.Quantity, req.TotalPrice); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{"message": newCartId})
		}
		return err
	}
	// update existing item in cart
	if err := h.db.QueryRow("UPDATE cart_items SET quantity = $1, total_price = $2 where product_id = $3 RETURNING id", req.Quantity+item.Quantity, req.TotalPrice+item.TotalPrice, req.ProductId).Scan(&cartId); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": cartId})
}

func (h *CartHandler) RemoveFromCart(c echo.Context) error {
	cartId := c.Param("id")
	if _, err := h.db.Exec("DELETE FROM cart_items where id = $1", cartId); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": cartId})
}

func (h *CartHandler) ClearCart(c echo.Context) error {
	userId := c.Get("userId").(string)
	if _, err := h.db.Exec("DELETE FROM cart_items where user_id = $1", userId); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
