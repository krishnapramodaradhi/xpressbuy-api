package handler

import (
	"database/sql"
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	db *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) FetchProducts(c echo.Context) error {
	rows, err := h.db.Query("SELECT p.id, p.title, p.short_description, p.description, p.price, p.quantity, p.image_url, c.id, c.title FROM products p, categories c where p.category = c.id")
	if err != nil {
		return err
	}
	defer rows.Close()

	products := make([]entity.Product, 0, 26)
	for rows.Next() {
		var p entity.Product
		if err = rows.Scan(&p.Id, &p.Title, &p.ShortDescription, &p.Description, &p.Price, &p.Quantity, &p.ImageUrl, &p.Category.Id, &p.Category.Title); err != nil {
			return err
		}
		products = append(products, p)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FetchProductById(c echo.Context) error {
	id := c.Param("id")
	row := h.db.QueryRow("SELECT p.id, p.title, p.short_description, p.description, p.price, p.quantity, p.image_url, c.id, c.title FROM products p, categories c where p.category = c.id and p.id = $1", id)
	var p entity.Product
	if err := row.Scan(&p.Id, &p.Title, &p.ShortDescription, &p.Description, &p.Price, &p.Quantity, &p.ImageUrl, &p.Category.Id, &p.Category.Title); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
