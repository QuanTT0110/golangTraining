package raw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Staff struct {
		ID           primitive.ObjectID `bson:"_id" form:"_id"`
		Name         string             `bson:"name" form:"name"`
		Email        string             `bson:"email" form:"email"`
		Password     string             `bson:"password,omitempty" form:"password"`
		DepartmentID primitive.ObjectID `bson:"department" form:"department"`
	}
)
