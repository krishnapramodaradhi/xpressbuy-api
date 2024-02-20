package middleware

import (
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util"
	customerror "github.com/krishnapramodaradhi/xpressbuy-api/internal/util/customError"
	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header["Authorization"]
		c.Logger().Debug(token)
		if len(token) == 0 {
			return customerror.New(http.StatusUnauthorized, "Auth Token not found")
		}
		payload, err := util.ValidateToken(token[0])
		if err != nil {
			return customerror.New(http.StatusUnauthorized, err.Error())
		}
		c.Set("userId", payload)
		return next(c)
	}
}
