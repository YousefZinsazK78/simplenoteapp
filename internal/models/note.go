package models

import "time"

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateNoteParams struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
