package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Created_at time.Time `json:"created_at"`
}

type UpdateUserParams struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
