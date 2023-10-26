package database

import (
	"context"
	"notegin/internal/models"
)

type UserStorer interface {
	InsertUser(context.Context, models.User) error
	ViewUsers(context.Context) ([]models.User, error)
	ViewUserByID(context.Context, int) (*models.User, error)
	ViewUserByUsername(context.Context, string) (*models.User, error)
	UpdateUser(context.Context, models.UpdateUserParams) error
	DeleteUser(context.Context, int) error
}

type userstore struct {
	Database
}

func NewUserStore(db Database) userstore {
	return userstore{
		Database: db,
	}
}

func (u userstore) InsertUser(ctx context.Context, user models.User) error {
	stmt, err := u.db.PrepareContext(ctx, "INSERT INTO user_tbl(username, password, email) values ($1,$2,$3);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u userstore) ViewUsers(c context.Context) ([]models.User, error) {
	query, err := u.db.QueryContext(c, "SELECT * FROM user_tbl;")
	if err != nil {
		return nil, err
	}
	defer query.Close()
	var userModel []models.User
	for query.Next() {
		var userm models.User
		if err := query.Scan(&userm.ID, &userm.Username, &userm.Password, &userm.Email, &userm.Created_at); err != nil {
			return nil, err
		}
		userModel = append(userModel, userm)
	}
	return userModel, nil
}

func (u userstore) ViewUserByID(c context.Context, id int) (*models.User, error) {
	row := u.db.QueryRowContext(c, "SELECT * FROM user_tbl WHERE id=$1", id)
	var userm models.User
	if err := row.Scan(&userm.ID, &userm.Username, &userm.Password, &userm.Email, &userm.Created_at); err != nil {
		return nil, err
	}
	return &userm, nil
}

func (u userstore) ViewUsersByUsername(c context.Context, username string) (*models.User, error) {
	row := u.db.QueryRowContext(c, "SELECT * FROM user_tbl WHERE username=$1", username)
	var userm models.User
	if err := row.Scan(&userm.ID, &userm.Username, &userm.Password, &userm.Email, &userm.Created_at); err != nil {
		return nil, err
	}
	return &userm, nil
}

func (u userstore) UpdateUser(ctx context.Context, user models.UpdateUserParams) error {
	stmt, err := u.db.PrepareContext(ctx, "UPDATE USER SET email=$1 WHERE id=$2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u userstore) DeleteUser(ctx context.Context, id int) error {
	stmt, err := u.db.PrepareContext(ctx, "DELETE FROM user_tbl WHERE id = $1")
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
