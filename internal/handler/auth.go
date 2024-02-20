package handler

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/entity"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util/constants"
	customerror "github.com/krishnapramodaradhi/xpressbuy-api/internal/util/customError"
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
		return customerror.New(http.StatusInternalServerError, err.Error())
	}

	hashedPassword, err := u.HashPassword(u.Password)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}

	u.Id = uuid.NewString()
	u.Password = hashedPassword
	result, err := h.db.Exec(constants.CREATE_USER, u.Id, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Info(result.LastInsertId())

	token, err := util.GenerateToken(u.Id)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) Login(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	var user entity.User
	row := h.db.QueryRow(constants.FIND_USER_EMAIL, u.Email)
	if err := row.Scan(&user.Id, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return customerror.New(http.StatusUnauthorized, "Auth Failed")
		}
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	err := user.ComparePassword(u.Password)
	if err != nil {
		return err
	}

	token, err := util.GenerateToken(user.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
