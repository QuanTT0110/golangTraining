package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	staffs      = "staffs"
	departments = "departments"
)

func StaffCol() *mongo.Collection {
	return db.Collection(staffs)
}

func DepartmentCol() *mongo.Collection {
	return db.Collection(departments)
}
