package postgres

import (
	"database/sql"
	"errors"

	"milliy/model"
	"milliy/storage"
)

type UsersRepo struct {
	Db *sql.DB
}

func NewUsersRepo(db *sql.DB) storage.UserStorage {
	return &UsersRepo{Db: db}
}

func (r *UsersRepo) CheckPassword(login, password string) (bool, error) {
	var hashedPassword string
	query := "SELECT password_hash FROM users WHERE login = $1"
	err := r.Db.QueryRow(query, login).Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	checkQuery := "SELECT crypt($1, password_hash) = password_hash FROM users WHERE login = $2"
	var isValid bool
	err = r.Db.QueryRow(checkQuery, password, login).Scan(&isValid)
	if err != nil {
		return false, err
	}

	return isValid, nil
}

func (r *UsersRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	query := `SELECT id, login FROM users WHERE id = $1`

	err := r.Db.QueryRow(query, id).Scan(&user.ID, &user.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
