package route

import "github.com/labstack/echo/v4"

func Route(e *echo.Echo) {
	Staff(e)
	Department(e)
}
