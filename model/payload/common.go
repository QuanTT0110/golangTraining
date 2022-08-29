package payload

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	LoginPayload struct {
		Email    string `json:"email" form:"email" bson:"email"`
		Password string `json:"password" form:"password" bson:"password"`
	}

	JwtCustomClaim struct {
		ID    primitive.ObjectID `json:"_id" form:"_id"`
		Name  string             `json:"name" form:"name"`
		Email string             `json:"admin" form:"admin"`
		jwt.StandardClaims
	}
)

func (payload LoginPayload) ValidateLogin() error {
	return validation.ValidateStruct(&payload,
		validation.Field(&payload.Email, validation.Required, is.Email),
		validation.Field(&payload.Password, validation.Required),
	)
}
