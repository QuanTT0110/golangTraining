package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"quanlyhoso/database"
	"quanlyhoso/model/query"
	"quanlyhoso/model/raw"
)

func CreateDepartment(ctx context.Context, department raw.Department) (raw.Department, error) {
	var (
		departmentCol = database.DepartmentCol()
	)

	_, err := departmentCol.InsertOne(ctx, department)

	return department, err
}

func UpdateDepartment(ctx context.Context, id primitive.ObjectID, department raw.Department) (raw.Department, error) {
	var (
		departmentCol = database.DepartmentCol()
		filter        = bson.D{{"_id", id}}
	)

	_, err := departmentCol.ReplaceOne(ctx, filter, department)

	return department, err
}

func DeleteDepartment(ctx context.Context, id primitive.ObjectID) error {
	var (
		departmentCol = database.DepartmentCol()
		filter        = bson.D{{"_id", id}}
	)

	_, err := departmentCol.DeleteOne(ctx, filter)

	return err
}

func GetDepartment(ctx context.Context, id primitive.ObjectID) (existingDepartment raw.Department, err error) {
	var (
		departmentCol = database.DepartmentCol()
		filter        = bson.D{{"_id", id}}
	)

	err = departmentCol.FindOne(ctx, filter).Decode(&existingDepartment)

	return existingDepartment, err
}

func GetAllDepartment(ctx context.Context, query query.DepartmentFindAllQuery) (departments []raw.Department, err error) {
	var departmentCol = database.DepartmentCol()
	var filter = bson.M{}

	if len(query.Keyword) > 0 {
		filter["name"] = bson.M{
			"$regex": query.Keyword,
		}
	}

	opts := options.Find().SetLimit(query.Limit).SetSkip((query.Page - 1) * query.Limit)
	result, err := departmentCol.Find(ctx, filter, opts)
	if err != nil {
		return departments, err
	}
	for result.Next(context.Background()) {
		var department raw.Department
		err = result.Decode(&department)
		if err != nil {
			log.Fatal(err)
		}
		departments = append(departments, department)
	}

	return departments, err
}
