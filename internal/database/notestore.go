package database

import (
	"context"
	"notegin/internal/models"
)

type NoteStorer interface {
	Insert(context.Context, models.Note) error
	GetAll(context.Context) ([]models.Note, error)
	GetByID(context.Context, int) (*models.Note, error)
	GetByTitle(context.Context, string) (*models.Note, error)
	Update(context.Context, models.UpdateNoteParams) error
	DeleteByID(context.Context, int) error
}

type noteStore struct {
	Database
}

func NewNoteStore(db Database) noteStore {
	return noteStore{
		Database: db,
	}
}

func (n noteStore) Insert(ctx context.Context, noteModel models.Note) error {
	stmt, err := n.db.PrepareContext(ctx, "INSERT INTO note_tbl(title, body,user_id) values ($1,$2,$3);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, noteModel.Title, noteModel.Body, noteModel.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (n noteStore) GetAll(ctx context.Context) ([]models.Note, error) {
	query, err := n.db.QueryContext(ctx, "SELECT * FROM note_tbl;")
	if err != nil {
		return nil, err
	}
	defer query.Close()
	var noteModels []models.Note
	for query.Next() {
		var notemodel models.Note
		if err := query.Scan(&notemodel.ID, &notemodel.Title, &notemodel.Body, &notemodel.UserID, &notemodel.CreatedAt); err != nil {
			return nil, err
		}
		noteModels = append(noteModels, notemodel)
	}
	return noteModels, nil
}

func (n noteStore) GetByID(ctx context.Context, id int) (*models.Note, error) {
	row := n.db.QueryRowContext(ctx, "SELECT * FROM note_tbl WHERE id=$1", id)
	var notemodel models.Note
	if err := row.Scan(&notemodel.ID, &notemodel.Title, &notemodel.Body, &notemodel.UserID, &notemodel.CreatedAt); err != nil {
		return nil, err
	}
	return &notemodel, nil
}

func (n noteStore) GetByTitle(ctx context.Context, title string) (*models.Note, error) {
	row := n.db.QueryRowContext(ctx, "SELECT * FROM note_tbl WHERE title LIKE '%$1%'", title)
	var notemodel models.Note
	if err := row.Scan(&notemodel.ID, &notemodel.Title, &notemodel.Body, &notemodel.UserID, &notemodel.CreatedAt); err != nil {
		return nil, err
	}
	return &notemodel, nil
}

func (u noteStore) Update(ctx context.Context, note models.UpdateNoteParams) error {
	stmt, err := u.db.PrepareContext(ctx, "UPDATE note_tbl SET title=$1 WHERE id=$2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, note.Title, note.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u noteStore) DeleteByID(ctx context.Context, id int) error {
	stmt, err := u.db.PrepareContext(ctx, "DELETE FROM note_tbl WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
