package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
)

type CrudsRepo struct {
	Db *sql.DB
}

func NewCrudsRepo(db *sql.DB) *CrudsRepo {
	return &CrudsRepo{Db: db}
}

func (r *CrudsRepo) CreateTwit(req *model.CreateTwitRequest) (string, error) {
	query := "INSERT INTO twit (user_id, texts, readers_count) VALUES ($1, $2, 0) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.UserID, req.Texts).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetTwitByID fetches a twit by its ID.
func (r *CrudsRepo) GetTwitByID(id string) (*model.Twit, error) {
	query := "SELECT id, user_id, texts, readers_count FROM twit WHERE id = $1 AND deleted_at = 0"
	var twit model.Twit
	err := r.Db.QueryRow(query, id).Scan(&twit.ID, &twit.UserID, &twit.Texts, &twit.ReadersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &twit, nil
}

// GetTwitsByUserID fetches all twits associated with a specific user ID.
func (r *CrudsRepo) GetTwitsByUserID(userID string) ([]*model.Twit, error) {
	query := "SELECT id, user_id, texts, readers_count FROM twit WHERE user_id = $1 AND deleted_at = 0"
	rows, err := r.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var twits []*model.Twit
	for rows.Next() {
		var twit model.Twit
		if err := rows.Scan(&twit.ID, &twit.UserID, &twit.Texts, &twit.ReadersCount); err != nil {
			return nil, err
		}
		twits = append(twits, &twit)
	}
	return twits, nil
}

// DeleteTwit soft-deletes a twit by its ID.
func (r *CrudsRepo) DeleteTwit(id string) error {
	query := "UPDATE twit SET deleted_at = EXTRACT(EPOCH FROM now()) WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// AddReadersCount increments the readers_count of a twit by 1.
func (r *CrudsRepo) AddReadersCount(id string) error {
	query := "UPDATE twit SET readers_count = readers_count + 1 WHERE id = $1"
	result, err := r.Db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("twit not found or already deleted")
	}

	return nil
}
