package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	hashedPassword, err := u.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Id = uuid.NewString()
	u.Password = hashedPassword
	result, err := h.db.Exec("INSERT INTO users (id, first_name, last_name, email, password) values ($1, $2, $3, $4, $5)", u.Id, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return err
	}
	c.Logger().Info(result.LastInsertId())

	token, err := util.FetchToken(u.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) Login(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	var user entity.User
	row := h.db.QueryRow("SELECT id, password FROM users where email = $1", u.Email)
	if err := row.Scan(&user.Id, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Auth Failed")
		}
		return err
	}
	err := user.ComparePassword(u.Password)
	if err != nil {
		return err
	}

	token, err := util.FetchToken(user.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
