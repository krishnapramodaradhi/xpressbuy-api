package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util/constants"
	customerror "github.com/krishnapramodaradhi/xpressbuy-api/internal/util/customError"
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
	rows, err := h.db.Query(constants.FETCH_CART, userId)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item entity.CartItem
		if err = rows.Scan(&item.Id, &item.Quantity, &item.TotalPrice, &item.Product.Id, &item.Product.Title, &item.Product.ImageUrl); err != nil {
			return customerror.New(http.StatusInternalServerError, err.Error())
		}
		cartItems = append(cartItems, item)
		totalPrice += item.TotalPrice
	}
	return c.JSON(http.StatusOK, map[string]any{"items": cartItems, "totalPrice": totalPrice})
}

func (h *CartHandler) AddItemToCart(c echo.Context) error {
	var req entity.CartItemRequest
	if err := c.Bind(&req); err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	var item entity.CartItem
	var cartId string
	row := h.db.QueryRow(constants.FIND_CART_ITEM, req.ProductId)
	if err := row.Scan(&item.Quantity, &item.TotalPrice); err != nil {
		if err == sql.ErrNoRows {
			// add item to cart
			newCartId := uuid.NewString()
			userId := c.Get("userId").(string)
			if _, err := h.db.Exec(constants.ADD_TO_CART, newCartId, userId, req.ProductId, req.Quantity, req.TotalPrice); err != nil {
				return customerror.New(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, map[string]string{"message": newCartId})
		}
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	// update existing item in cart
	if err := h.db.QueryRow(constants.UPDATE_CART, req.Quantity+item.Quantity, req.TotalPrice+item.TotalPrice, req.ProductId).Scan(&cartId); err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": cartId})
}

func (h *CartHandler) RemoveFromCart(c echo.Context) error {
	cartId := c.Param("id")
	if _, err := h.db.Exec(constants.DELETE_CART_ITEM, cartId); err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": cartId})
}

func (h *CartHandler) ClearCart(c echo.Context) error {
	userId := c.Get("userId").(string)
	if _, err := h.db.Exec(constants.CLEAR_CART, userId); err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
