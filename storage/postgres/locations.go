package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
	"milliy/storage"
)

type LocationsRepo struct {
	Db *sql.DB
}

func NewLocationsRepo(db *sql.DB) storage.LocationStorage {
	return &LocationsRepo{Db: db}
}

// Create inserts a new location record.
func (r *LocationsRepo) Create(req *model.CreateLocationRequest) (string, error) {
	query := "INSERT INTO locations (twit_id, lat, lon) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.TwitID, req.Lat, req.Lon).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// DeleteByID deletes a location by its ID.
func (r *LocationsRepo) DeleteByID(id string) error {
	query := "DELETE FROM locations WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// DeleteByTwitID deletes locations associated with a specific Twit ID.
func (r *LocationsRepo) DeleteByTwitID(twitID string) error {
	query := "DELETE FROM locations WHERE twit_id = $1"
	_, err := r.Db.Exec(query, twitID)
	return err
}

// GetByTwitID fetches locations by Twit ID.
func (r *LocationsRepo) GetByTwitID(twitID string) ([]model.Location, error) {
	query := "SELECT id, twit_id, lat, lon FROM locations WHERE twit_id = $1"
	rows, err := r.Db.Query(query, twitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []model.Location
	for rows.Next() {
		var location model.Location
		if err := rows.Scan(&location.ID, &location.TwitID, &location.Lat, &location.Lon); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}
	return locations, nil
}

// GetByID fetches a location by its ID.
func (r *LocationsRepo) GetByID(id string) (*model.Location, error) {
	query := "SELECT id, twit_id, lat, lon FROM locations WHERE id = $1"
	var location model.Location
	err := r.Db.QueryRow(query, id).Scan(&location.ID, &location.TwitID, &location.Lat, &location.Lon)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &location, nil
}

// UpdateByID updates a location by its ID.
func (r *LocationsRepo) UpdateByID(id string, req *model.CreateLocationRequest) error {
	query := "UPDATE locations SET lat = $1, lon = $2 WHERE id = $3"
	_, err := r.Db.Exec(query, req.Lat, req.Lon, id)
	return err
}
