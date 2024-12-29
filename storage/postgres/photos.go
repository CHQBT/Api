package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
	"milliy/storage"
)

type PhotosRepo struct {
	Db *sql.DB
}

func NewPhotosRepo(db *sql.DB) storage.PhotoStorage {
	return &PhotosRepo{Db: db}
}

// Create inserts a new photo record.
func (r *PhotosRepo) Create(req *model.CreatePhotoRequest) (string, error) {
	query := "INSERT INTO photos (twit_id, photo) VALUES ($1, $2) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.TwitID, req.Photo).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// DeleteByID deletes a photo by its ID.
func (r *PhotosRepo) DeleteByID(id string) error {
	query := "DELETE FROM photos WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// DeleteByTwitID deletes photos associated with a specific Twit ID.
func (r *PhotosRepo) DeleteByTwitID(twitID string) error {
	query := "DELETE FROM photos WHERE twit_id = $1"
	_, err := r.Db.Exec(query, twitID)
	return err
}

// GetByTwitID fetches photos by Twit ID.
func (r *PhotosRepo) GetByTwitID(twitID string) ([]model.Photo, error) {
	query := "SELECT id, twit_id, photo FROM photos WHERE twit_id = $1"
	rows, err := r.Db.Query(query, twitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []model.Photo
	for rows.Next() {
		var photo model.Photo
		if err := rows.Scan(&photo.ID, &photo.TwitID, &photo.Photo); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

// GetByID fetches a photo by its ID.
func (r *PhotosRepo) GetByID(id string) (*model.Photo, error) {
	query := "SELECT id, twit_id, photo FROM photos WHERE id = $1"
	var photo model.Photo
	err := r.Db.QueryRow(query, id).Scan(&photo.ID, &photo.TwitID, &photo.Photo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &photo, nil
}

// UpdateByID updates a photo by its ID.
func (r *PhotosRepo) UpdateByID(id string, req *model.CreatePhotoRequest) error {
	query := "UPDATE photos SET photo = $1 WHERE id = $2"
	_, err := r.Db.Exec(query, req.Photo, id)
	return err
}
