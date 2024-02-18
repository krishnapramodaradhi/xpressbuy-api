package middleware

import (
	"net/http"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util"
	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header["Authorization"]
		if len(token) == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Auth Token not found")
		}
		payload, err := util.ValidateToken(token[0])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		c.Set("userId", payload)
		return next(c)
	}
}
