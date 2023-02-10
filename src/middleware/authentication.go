package middleware

import (
	"todoapp/src/abstraction"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*abstraction.Context)
		cc.Auth = &abstraction.AuthContext{
			ID:    0,
			Name:  "system",
		}

		return next(cc)
	}
}
