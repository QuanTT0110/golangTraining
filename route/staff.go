package route

import (
	"github.com/labstack/echo/v4"
	"quanlyhoso/controller"
	"quanlyhoso/validation"
)

func Staff(e *echo.Echo) {
	staffs := e.Group("/staffs")

	//staffs.Use(middleware.JWTWithConfig(permission.JWTConfig))
	staffs.POST("", controller.CreateStaff, validation.CreateStaff)
	staffs.PUT("/:id", controller.UpdateStaff, validation.ValidateStaffID, validation.CheckStaffExistedByID, validation.CreateStaff)
	staffs.DELETE("/:id", controller.DeleteStaff, validation.ValidateStaffID, validation.CheckStaffExistedByID)
	staffs.GET("/:id", controller.GetStaff, validation.ValidateStaffID, validation.CheckStaffExistedByID)
	staffs.GET("", controller.GetAllStaff, validation.StaffQuery)

}
