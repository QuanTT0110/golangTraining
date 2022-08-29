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

func CreateStaff(ctx context.Context, staff raw.Staff) (raw.Staff, error) {

	var (
		staffCol = database.StaffCol()
	)

	_, err := staffCol.InsertOne(ctx, staff)

	return staff, err
}

func UpdateStaff(ctx context.Context, id primitive.ObjectID, staff raw.Staff) (raw.Staff, error) {
	var (
		staffCol = database.StaffCol()
		filter   = bson.D{{"_id", id}}
	)

	_, err := staffCol.ReplaceOne(ctx, filter, staff)

	return staff, err
}

func DeleteStaff(ctx context.Context, id primitive.ObjectID) error {
	var (
		staffCol = database.StaffCol()
		filter   = bson.D{{"_id", id}}
	)

	_, err := staffCol.DeleteOne(ctx, filter)

	return err
}

func GetStaff(ctx context.Context, id primitive.ObjectID) (existingStaff raw.Staff, err error) {
	var (
		staffCol = database.StaffCol()
		filter   = bson.D{{"_id", id}}
	)

	err = staffCol.FindOne(ctx, filter).Decode(&existingStaff)

	return existingStaff, err
}

func GetAllStaff(ctx context.Context, query query.StaffFindAllQuery) (staffs []raw.Staff, err error) {
	var staffCol = database.StaffCol()
	var filter = bson.M{}

	if !query.Department.IsZero() {
		filter["departmentId"] = query.Department
	}

	if len(query.Keyword) > 0 {
		filter["name"] = bson.M{
			"$regex": query.Keyword,
		}
	}

	opts := options.Find().SetLimit(query.Limit).SetSkip((query.Page - 1) * query.Limit)
	result, err := staffCol.Find(ctx, filter, opts)
	if err != nil {
		return staffs, err
	}
	for result.Next(context.Background()) {
		var staff raw.Staff
		err = result.Decode(&staff)
		if err != nil {
			log.Fatal(err)
		}
		staffs = append(staffs, staff)
	}

	return staffs, err
}

func FindByEmail(ctx context.Context, email string) (raw.Staff, error) {
	var (
		staffCol      = database.StaffCol()
		filter        = bson.D{{"email", email}}
		existingStaff raw.Staff
	)

	err := staffCol.FindOne(ctx, filter).Decode(&existingStaff)

	return existingStaff, err
}
