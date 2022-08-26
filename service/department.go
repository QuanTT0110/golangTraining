package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quanlyhoso/dao"
	"quanlyhoso/model/payload"
	"quanlyhoso/model/query"
	"quanlyhoso/model/response"
)

func CreateDepartment(ctx context.Context, payload payload.DepartmentCreatePayLoad) (res response.DepartmentResponse, err error) {
	department := payload.ConvertToBSON()
	createdDepartment, err := dao.CreateDepartment(ctx, department)
	if err != nil {
		err = errors.New("Create department error")
		return res, err
	}

	res = response.DepartmentResponse{
		ID:      createdDepartment.ID.Hex(),
		Name:    createdDepartment.Name,
		Address: createdDepartment.Address,
	}

	return res, err
}

func UpdateDepartment(ctx context.Context, id primitive.ObjectID, payload payload.DepartmentCreatePayLoad) (res response.DepartmentResponse, err error) {
	department := payload.ConvertToBSON()
	updatedDepartment, err := dao.UpdateDepartment(ctx, id, department)
	if err != nil {
		return res, errors.New("Update department error")
	}

	res = response.DepartmentResponse{
		ID:      updatedDepartment.ID.Hex(),
		Name:    updatedDepartment.Name,
		Address: updatedDepartment.Address,
	}

	return res, err
}

func DeleteDepartment(ctx context.Context, id primitive.ObjectID) error {
	err := dao.DeleteDepartment(ctx, id)
	if err != nil {
		err = errors.New("Delete department error")
		return err
	}

	return nil
}

func GetDepartment(ctx context.Context, id primitive.ObjectID) (res response.DepartmentResponse, err error) {
	department, err := dao.GetDepartment(ctx, id)
	if err != nil {
		return res, errors.New("Department not found")
	}

	res = response.DepartmentResponse{
		ID:      department.ID.Hex(),
		Name:    department.Name,
		Address: department.Address,
	}

	return res, err
}

func GetAllDepartment(ctx context.Context, query query.DepartmentFindAllQuery) (res []response.DepartmentResponse, err error) {
	rs, err := dao.GetAllDepartment(ctx, query)
	if err != nil {
		return res, errors.New("Can not get list of department")
	}

	for _, value := range rs {

		res = append(res, response.DepartmentResponse{
			ID:      value.ID.Hex(),
			Name:    value.Name,
			Address: value.Address,
		})
	}

	return res, err
}
