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
	u := new(entity.UserRequest)
	if err := c.Bind(u); err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(u); err != nil {
		c.Logger().Error("An error occured with the request validation", err)
		return customerror.New(http.StatusBadRequest, err.Error())
	}
	newUser := entity.NewUser(uuid.NewString(), u.FirstName, u.LastName, u.Email)
	hashedPassword, err := newUser.HashPassword(u.Password)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}

	newUser.Password = hashedPassword
	_, err = h.db.Exec(constants.CREATE_USER, newUser.Id, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, err.Error())
	}

	token, err := util.GenerateToken(newUser.Id)
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
			return customerror.New(http.StatusBadRequest, "Auth Failed")
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
