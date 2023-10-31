package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserParamsForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type JwtUserClaims struct {
	Userid int `json:"user_id"`
	jwt.RegisteredClaims
}

type UpdateUserParams struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
