package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserJWT struct{
	Email string
	Role string
}

func (u *UserJWT)  ToJwtClaims() jwt.Claims{
	return jwt.MapClaims{
		"email": u.Email,
		"role" : u.Role,
		"exp" : time.Now().Add(24 * time.Hour).Unix(),
	}
}