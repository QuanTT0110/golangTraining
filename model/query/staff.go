package query

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffFindAllQuery struct {
	Limit      int64              `query:"limit"`
	Page       int64              `query:"page"`
	Keyword    string             `query:"keyword"`
	Department primitive.ObjectID `query:"department,omitempty"`
}

func (query StaffFindAllQuery) ValidateStaffQuery() error {
	return validation.ValidateStruct(&query,
		validation.Field(
			&query.Limit,
			validation.Required,
			validation.Min(1).Error("Limit is > 0"),
			validation.Max(20).Error("Limit is < 20"),
		),
		validation.Field(
			&query.Page,
			validation.Required,
			validation.Min(1).Error("Page is > 0"),
		),
		validation.Field(
			&query.Department,
			validation.Each(is.MongoID),
		),
	)
}
