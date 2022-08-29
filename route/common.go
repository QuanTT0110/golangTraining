package route

import (
	"github.com/labstack/echo/v4"
	"quanlyhoso/controller"
	"quanlyhoso/validation"
)

func Common(e *echo.Echo) {
	common := e.Group("/")

	common.POST("login", controller.Login, validation.Login)

}
