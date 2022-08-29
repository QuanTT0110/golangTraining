package validation

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"quanlyhoso/dao"
	"quanlyhoso/model/payload"
	"quanlyhoso/model/query"
)

func CreateDepartment(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.DepartmentCreatePayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateCreateDepartment()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func ValidateDepartmentID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id                = c.Param("id")
			departmentID, err = ValidateObjectID(id)
		)

		if err != nil {
			return c.JSON(http.StatusBadRequest, error.Error(errors.New("Department ID is not valid")))
		}

		c.Set("departmentID", departmentID)
		return next(c)
	}
}

func CheckDepartmentExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var (
			departmentID = c.Get("departmentID").(primitive.ObjectID)
		)

		department, _ := dao.GetDepartment(ctx, departmentID)
		if department.ID.Hex() == "" {
			return c.JSON(http.StatusNoContent, error.Error(errors.New("Department not found")))
		}

		c.Set("department", department)
		return next(c)
	}
}

func DepartmentQuery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query query.DepartmentFindAllQuery
		)

		err := c.Bind(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, error.Error(errors.New("Invalid input")))
		}

		err = query.ValidateDepartmentQuery()
		if err != nil {
			c.JSON(http.StatusBadRequest, error.Error(errors.New("Invalid input")))
		}

		c.Set("query", query)
		return next(c)
	}
}
