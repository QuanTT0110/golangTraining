package route

import (
	"github.com/labstack/echo/v4"
	"quanlyhoso/controller"
	"quanlyhoso/validation"
)

func Department(e *echo.Echo) {
	departments := e.Group("/departments")

	//departments.Use(middleware.JWTWithConfig(permission.JWTConfig))
	departments.POST("", controller.CreateDepartment, validation.CreateDepartment)
	departments.PUT("/:id", controller.UpdateDepartment, validation.ValidateDepartmentID, validation.CheckDepartmentExistedByID, validation.CreateDepartment)
	departments.DELETE("/:id", controller.DeleteDepartment, validation.ValidateDepartmentID, validation.CheckDepartmentExistedByID)
	departments.GET("/:id", controller.GetDepartment, validation.ValidateDepartmentID, validation.CheckDepartmentExistedByID)
	departments.GET("", controller.GetAllDepartment, validation.DepartmentQuery)

}
