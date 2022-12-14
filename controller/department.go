package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"quanlyhoso/model/payload"
	"quanlyhoso/model/query"
	"quanlyhoso/service"
)

func CreateDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var payload, ok = c.Get("payload").(payload.DepartmentCreatePayLoad)
	if !ok {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Bad Request")))
	}

	createdDepartment, err := service.CreateDepartment(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, createdDepartment)
}

func UpdateDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	var payload, ok = c.Get("payload").(payload.DepartmentCreatePayLoad)
	if !ok {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Bad Request")))
	}

	updatedDepartment, err := service.UpdateDepartment(ctx, ID, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedDepartment)
}

func DeleteDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	err := service.DeleteDepartment(ctx, ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, error.Error(errors.New("Delete Department successfully")))
}

func GetDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	department, err := service.GetDepartment(ctx, ID)

	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}

	return c.JSON(http.StatusOK, department)
}

func GetAllDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	query, ok := c.Get("query").(query.DepartmentFindAllQuery)
	if !ok {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Bad Request")))
	}

	pagedDepartment, err := service.GetAllDepartment(ctx, query)
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}

	return c.JSON(http.StatusOK, pagedDepartment)
}
