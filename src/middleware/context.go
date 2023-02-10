package middleware

import (
	"todoapp/src/abstraction"

	"github.com/labstack/echo/v4"
)

func Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &abstraction.Context{
			Context: c,
		}
		return next(cc)
	}
}