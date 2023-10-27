package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Created_at time.Time `json:"created_at"`
}

type JwtUserClaims struct {
	Userid int `json:"user_id"`
	jwt.RegisteredClaims
}

type UpdateUserParams struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
