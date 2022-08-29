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
		fmt.Println(payload)
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
			return c.JSON(http.StatusNoContent, error.Error(errors.New("Staff not found")))
		}

		c.Set("staff", staff)
		return next(c)
	}
}

func CheckEmailExisted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var (
			email = c.Param("email")
		)

		staff, _ := dao.FindByEmail(ctx, email)
		fmt.Println(staff.ID)
		if !staff.ID.IsZero() {
			return c.JSON(http.StatusConflict, error.Error(errors.New("Email is existing")))
		}

		return next(c)
	}
}
