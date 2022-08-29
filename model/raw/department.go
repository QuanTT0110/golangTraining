package raw

import "go.mongodb.org/mongo-driver/bson/primitive"

type Department struct {
	ID      primitive.ObjectID `bson:"_id" form:"_id"`
	Name    string             `bson:"name" form:"name"`
	Address string             `bson:"address" form:"address"`
}
