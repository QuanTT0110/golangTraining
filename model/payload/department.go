package payload

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quanlyhoso/model/raw"
)

type DepartmentCreatePayLoad struct {
	ID      string `json:"_id" form:"_id"`
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

func (payload DepartmentCreatePayLoad) ConvertToBSON() raw.Department {
	result := raw.Department{
		ID:      primitive.NewObjectID(),
		Name:    payload.Name,
		Address: payload.Address,
	}
	return result
}

func (payload DepartmentCreatePayLoad) ValidateCreateDepartment() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Name,
			validation.Required.Error("Name is required"),
			validation.Length(3, 30).Error("Name is length: 3 -> 30"),
			is.Alpha.Error("Name is alphabet"),
		),
		validation.Field(
			&payload.Address,
			validation.Required.Error("Address is required"),
			validation.Length(10, 300).Error("Name is length: 3 -> 30"),
		),
	)
}
