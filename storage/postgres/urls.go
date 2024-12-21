package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
)

type URLsRepo struct {
	Db *sql.DB
}

func NewURLsRepo(db *sql.DB) *URLsRepo {
	return &URLsRepo{Db: db}
}

// Create inserts a new URL record.
func (r *URLsRepo) Create(req *model.CreateURLRequest) (string, error) {
	query := "INSERT INTO urls (twit_id, url) VALUES ($1, $2) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.TwitID, req.URL).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// DeleteByID deletes a URL by its ID.
func (r *URLsRepo) DeleteByID(id string) error {
	query := "DELETE FROM urls WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// DeleteByTwitID deletes URLs associated with a specific Twit ID.
func (r *URLsRepo) DeleteByTwitID(twitID string) error {
	query := "DELETE FROM urls WHERE twit_id = $1"
	_, err := r.Db.Exec(query, twitID)
	return err
}

// GetByTwitID fetches URLs by Twit ID.
func (r *URLsRepo) GetByTwitID(twitID string) ([]model.URL, error) {
	query := "SELECT id, twit_id, url FROM urls WHERE twit_id = $1"
	rows, err := r.Db.Query(query, twitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []model.URL
	for rows.Next() {
		var url model.URL
		if err := rows.Scan(&url.ID, &url.TwitID, &url.URL); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

// GetByID fetches a URL by its ID.
func (r *URLsRepo) GetByID(id string) (*model.URL, error) {
	query := "SELECT id, twit_id, url FROM urls WHERE id = $1"
	var url model.URL
	err := r.Db.QueryRow(query, id).Scan(&url.ID, &url.TwitID, &url.URL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

// UpdateByID updates a URL by its ID.
func (r *URLsRepo) UpdateByID(id string, req *model.CreateURLRequest) error {
	query := "UPDATE urls SET url = $1 WHERE id = $2"
	_, err := r.Db.Exec(query, req.URL, id)
	return err
}
