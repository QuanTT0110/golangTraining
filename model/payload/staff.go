package payload

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quanlyhoso/model/raw"
)

type StaffCreatePayLoad struct {
	ID           string `json:"_id" bson:"id"`
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password,omitempty" form:"password"`
	DepartmentID string `json:"department" form:"department"`
}

//type StaffLoginPayLoad struct {
//	Email    string `json:"email" form:"email"`
//	Password string `json:"password" form:"password"`
//}

func (payload StaffCreatePayLoad) ValidateCreateStaff() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Name,
			validation.Required.Error("Name is required"),
			validation.Length(3, 30).Error("Name is length: 3 -> 30"),
			is.Alpha.Error("Name is alphabet"),
		),
		validation.Field(
			&payload.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Invalid email"),
		),
		validation.Field(
			&payload.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 30).Error("Password is length: 8 -> 30"),
		),
	)
}

func (payload StaffCreatePayLoad) ConvertToBSON() raw.Staff {
	departmentID, _ := primitive.ObjectIDFromHex(payload.DepartmentID)
	result := raw.Staff{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Email:        payload.Email,
		Password:     payload.Password,
		DepartmentID: departmentID,
	}
	return result
}
