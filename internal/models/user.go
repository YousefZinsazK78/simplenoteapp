package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

type UpdateUserParams struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
