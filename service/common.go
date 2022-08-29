package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quanlyhoso/config"
	"quanlyhoso/model/payload"
	"quanlyhoso/model/response"
	"quanlyhoso/util"
)

func Login(ctx context.Context, login payload.LoginPayload) (res response.LoginResponse, err error) {
	staff, err := GetStaffByEmail(ctx, login.Email)

	staffID, _ := primitive.ObjectIDFromHex(staff.ID)

	if err != nil {
		return res, err
	}
	if !util.ComparePassword(login.Password, staff.Password) {
		return res, errors.New("Wrong password")
	}

	var claims = payload.JwtCustomClaim{
		ID:    staffID,
		Name:  staff.Name,
		Email: staff.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: util.CreateExpiryByHours(24),
		},
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var encode, err2 = token.SignedString([]byte(config.Env.SigningKey))
	if err2 != nil {
		return res, errors.New("Fail")
	}

	res = response.LoginResponse{
		Token: encode,
		ID:    staffID.Hex(),
		Email: staff.Email,
		Name:  staff.Name,
	}

	return res, err
}
