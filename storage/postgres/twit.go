package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
	"milliy/storage"
)

type TwitRepo struct {
	Db *sql.DB
}

func NewTwitRepo(db *sql.DB) storage.TwitStorage {
	return &TwitRepo{Db: db}
}

func (r *TwitRepo) CreateTwit(req *model.CreateTwitRequest) (string, error) {
	query := "INSERT INTO twit (user_id, title, texts, readers_count) VALUES ($1, $2, 0) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.UserID, req.Title, req.Texts).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetTwitByID fetches a twit by its ID.
func (r *TwitRepo) GetTwitByID(id string) (*model.Twit, error) {
	query := "SELECT id, user_id, title, texts, readers_count FROM twit WHERE id = $1 AND deleted_at = 0"
	var twit model.Twit
	err := r.Db.QueryRow(query, id).Scan(&twit.ID, &twit.UserID, &twit.Title, &twit.Texts, &twit.ReadersCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &twit, nil
}

// DeleteTwit soft-deletes a twit by its ID.
func (r *TwitRepo) DeleteTwit(id string) error {
	query := "UPDATE twit SET deleted_at = EXTRACT(EPOCH FROM now()) WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// AddReadersCount increments the readers_count of a twit by 1.
func (r *TwitRepo) AddReadersCount(id string) error {
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

func (p *TwitRepo) GetMostViewedTwit(limit int) ([]model.Twit, error) {
	var twits []model.Twit
	query := `SELECT * FROM twit ORDER BY readers_count DESC LIMIT $1`
	rows, err := p.Db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var twit model.Twit
		if err := rows.Scan(&twit.ID, &twit.UserID, &twit.Texts, &twit.Title, &twit.ReadersCount); err != nil {
			return nil, err
		}
		twits = append(twits, twit)
	}
	return twits, nil
}

func (p *TwitRepo) GetLatestTwits(limit int) ([]model.Twit, error) {
	var twits []model.Twit
	query := `SELECT * FROM twit ORDER BY created_at DESC LIMIT $1`
	rows, err := p.Db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var twit model.Twit
		if err := rows.Scan(&twit.ID, &twit.UserID, &twit.Texts, &twit.Title, &twit.ReadersCount); err != nil {
			return nil, err
		}
		twits = append(twits, twit)
	}
	return twits, nil
}

func (p *TwitRepo) SearchTwit(keyword string) ([]model.Twit, error) {
	var twits []model.Twit
	query := `SELECT * FROM twit WHERE title ILIKE '%' || $1 || '%' OR texts ILIKE '%' || $1 || '%'`
	rows, err := p.Db.Query(query, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var twit model.Twit
		if err := rows.Scan(&twit.ID, &twit.UserID, &twit.Texts, &twit.Title, &twit.ReadersCount); err != nil {
			return nil, err
		}
		twits = append(twits, twit)
	}
	return twits, nil
}
