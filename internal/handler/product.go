package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util/constants"
	customerror "github.com/krishnapramodaradhi/xpressbuy-api/internal/util/customError"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	db *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) FetchProducts(c echo.Context) error {
	rows, err := h.db.Query(constants.FETCH_ALL_PRODUCTS)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	products := make([]entity.Product, 0, 26)
	for rows.Next() {
		var p entity.Product
		if err = rows.Scan(&p.Id, &p.Title, &p.ShortDescription, &p.Description, &p.Price, &p.Quantity, &p.ImageUrl, &p.Category.Id, &p.Category.Title); err != nil {
			return customerror.New(http.StatusInternalServerError, err.Error())
		}
		products = append(products, p)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FetchProductById(c echo.Context) error {
	id := c.Param("id")
	row := h.db.QueryRow(constants.FETCH_PRODUCT_BY_ID, id)
	var p entity.Product
	if err := row.Scan(&p.Id, &p.Title, &p.ShortDescription, &p.Description, &p.Price, &p.Quantity, &p.ImageUrl, &p.Category.Id, &p.Category.Title); err != nil {
		if err == sql.ErrNoRows {
			return customerror.New(http.StatusNotFound, fmt.Sprintf("product with id %v is not found", id))
		}
		return customerror.New(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
