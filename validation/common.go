package validation

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"quanlyhoso/model/payload"
)

func Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.LoginPayload
		)

		c.Bind(&payload)

		err := payload.ValidateLogin()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}
