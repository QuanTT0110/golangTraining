package validation

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"quanlyhoso/dao"
	"quanlyhoso/model/payload"
	"quanlyhoso/model/query"
)

func CreateStaff(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.StaffCreatePayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateCreateStaff()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func StaffQuery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query query.StaffFindAllQuery
		)

		err := c.Bind(&query)
		if err != nil {
			c.JSON(http.StatusBadRequest, error.Error(errors.New("Invalid input")))
		}

		err = query.ValidateStaffQuery()
		if err != nil {
			c.JSON(http.StatusBadRequest, error.Error(errors.New("Invalid input")))
		}

		c.Set("query", query)
		return next(c)
	}
}

func ValidateObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
	}

	return objectID, err
}

func ValidateStaffID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id           = c.Param("id")
			staffID, err = ValidateObjectID(id)
		)

		if err != nil {
			return c.JSON(http.StatusBadRequest, error.Error(errors.New("Staff ID is not valid")))
		}

		c.Set("staffID", staffID)
		return next(c)
	}
}

func CheckStaffExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var (
			staffID = c.Get("staffID").(primitive.ObjectID)
		)

		staff, _ := dao.GetStaff(ctx, staffID)
		if staff.ID.IsZero() {
			return c.JSON(http.StatusNotFound, error.Error(errors.New("Staff not found")))
		}

		c.Set("staff", staff)
		return next(c)
	}
}
