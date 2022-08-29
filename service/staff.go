package service

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quanlyhoso/dao"
	"quanlyhoso/model/payload"
	"quanlyhoso/model/query"
	"quanlyhoso/model/response"
	"quanlyhoso/util"
)

func CreateStaff(ctx context.Context, payload payload.StaffCreatePayLoad) (res response.StaffResponse, err error) {
	departmentID, _ := primitive.ObjectIDFromHex(payload.DepartmentID)
	findDepartment, _ := GetDepartment(ctx, departmentID)
	fmt.Println(findDepartment)
	payload.DepartmentID = findDepartment.ID
	payload.Password, _ = util.HashPassword(payload.Password)

	staff := payload.ConvertToBSON()
	createdStaff, err := dao.CreateStaff(ctx, staff)
	if err != nil {
		err = errors.New("Create staff error")
		return res, err
	}

	res = response.StaffResponse{
		ID:       createdStaff.ID.Hex(),
		Name:     createdStaff.Name,
		Email:    createdStaff.Email,
		Password: "",
		Department: response.DepartmentResponse{
			ID:      findDepartment.ID,
			Name:    findDepartment.Name,
			Address: findDepartment.Address,
		},
	}

	return res, err
}

func UpdateStaff(ctx context.Context, id primitive.ObjectID, payload payload.StaffCreatePayLoad) (res response.StaffResponse, err error) {
	staff := payload.ConvertToBSON()
	updatedStaff, err := dao.UpdateStaff(ctx, id, staff)

	if err != nil {
		return res, errors.New("Update staff error")
	}

	res = response.StaffResponse{
		ID:       updatedStaff.ID.Hex(),
		Name:     updatedStaff.Name,
		Email:    updatedStaff.Email,
		Password: "",
		Department: response.DepartmentResponse{
			ID: updatedStaff.DepartmentID.Hex(),
		},
	}

	return res, err
}

func DeleteStaff(ctx context.Context, id primitive.ObjectID) error {
	err := dao.DeleteStaff(ctx, id)

	if err != nil {
		err = errors.New("Delete staff error")
		return err
	}
	return nil
}

func GetStaff(ctx context.Context, id primitive.ObjectID) (res response.StaffResponse, err error) {
	staff, err := dao.GetStaff(ctx, id)

	if err != nil {
		return res, errors.New("Staff not found")
	}

	res = response.StaffResponse{
		ID:    staff.ID.Hex(),
		Name:  staff.Name,
		Email: staff.Email,
		Department: response.DepartmentResponse{
			ID: staff.ID.Hex(),
		},
	}

	return res, err
}

func GetAllStaff(ctx context.Context, query query.StaffFindAllQuery) (res []response.StaffResponse, err error) {
	rs, err := dao.GetAllStaff(ctx, query)
	if err != nil {
		return res, errors.New("Can not get list of staff")
	}

	for _, value := range rs {
		department, _ := GetDepartment(ctx, value.DepartmentID)

		res = append(res, response.StaffResponse{
			ID:         value.ID.Hex(),
			Name:       value.Name,
			Email:      value.Email,
			Password:   value.Password,
			Department: department,
		})
	}

	return res, err
}

func GetStaffByEmail(ctx context.Context, email string) (res response.StaffResponse, err error) {
	staff, err := dao.FindByEmail(ctx, email)

	if err != nil {
		return res, errors.New("Staff not found")
	}

	res = response.StaffResponse{
		ID:       staff.ID.Hex(),
		Name:     staff.Name,
		Email:    staff.Email,
		Password: staff.Password,
		Department: response.DepartmentResponse{
			ID: staff.DepartmentID.Hex(),
		},
	}

	return res, err
}
