package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		AuthHeader := c.Request().Header.Get("Authorization")

		if strings.Contains(AuthHeader, "wrong") {
			return echo.ErrUnauthorized
		}
		return next(c)
	}

}
