package permission

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"quanlyhoso/config"
	"quanlyhoso/model/payload"
	"quanlyhoso/service"
	"strings"
)

var JWTConfig = middleware.JWTConfig{
	TokenLookup: "header:authorization",
	ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
			}
			return []byte(config.Env.SigningKey), nil
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			return nil, errors.New("Invalid token")
		}
		var splitAuth = strings.Split(auth, "Bearer ")
		token, err := jwt.Parse(splitAuth[1], keyFunc)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("Invalid token")
		}
		var account = payload.JwtCustomClaim{}
		jsonClaim, _ := json.Marshal(token.Claims)
		if err := json.Unmarshal(jsonClaim, &account); err != nil {
			return nil, errors.New("Fail")
		}
		var ctx = c.Request().Context()
		var user, _ = service.GetStaff(ctx, account.ID)
		return user, nil
	},
}
