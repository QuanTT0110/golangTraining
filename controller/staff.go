package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	payload "quanlyhoso/model/payload"
	"quanlyhoso/model/query"
	"quanlyhoso/service"
)

func CreateStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var payload, ok = c.Get("payload").(payload.StaffCreatePayLoad)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	createdStaff, err := service.CreateStaff(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, createdStaff)
}

func UpdateStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	var payload, ok = c.Get("payload").(payload.StaffCreatePayLoad)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	updatedStaff, err := service.UpdateStaff(ctx, ID, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, updatedStaff)
}

func DeleteStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	err := service.DeleteStaff(ctx, ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusNoContent, "Delete staff successfully")
}

func GetStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	staff, err := service.GetStaff(ctx, ID)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	return c.JSON(http.StatusOK, staff)
}

func GetAllStaff(c echo.Context) error {
	ctx := c.Request().Context()
	query, ok := c.Get("query").(query.StaffFindAllQuery)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	pagedStaff, err := service.GetAllStaff(ctx, query)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	return c.JSON(http.StatusOK, pagedStaff)
}
