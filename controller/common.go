package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	payload "quanlyhoso/model/payload"
	"quanlyhoso/service"
)

func Login(c echo.Context) error {
	ctx := c.Request().Context()
	payload := c.Get("payload").(payload.LoginPayload)
	fmt.Println(payload)
	result, err := service.Login(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
