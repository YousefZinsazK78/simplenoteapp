package models

import "time"

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type NoteForm struct {
	ID        int       `form:"id"`
	Title     string    `form:"title"`
	Body      string    `form:"body"`
	UserID    int       `form:"user_id"`
	CreatedAt time.Time `form:"created_at"`
}

type UpdateNoteParams struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
