package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
)

type MusicsRepo struct {
	Db *sql.DB
}

func NewMusicsRepo(db *sql.DB) *MusicsRepo {
	return &MusicsRepo{Db: db}
}

// Create inserts a new music record.
func (r *MusicsRepo) Create(req *model.CreateMusicRequest) (string, error) {
	query := "INSERT INTO musics (twit_id, mp3) VALUES ($1, $2) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.TwitID, req.MP3).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// DeleteByID deletes a music by its ID.
func (r *MusicsRepo) DeleteByID(id string) error {
	query := "DELETE FROM musics WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// DeleteByTwitID deletes musics associated with a specific Twit ID.
func (r *MusicsRepo) DeleteByTwitID(twitID string) error {
	query := "DELETE FROM musics WHERE twit_id = $1"
	_, err := r.Db.Exec(query, twitID)
	return err
}

// GetByTwitID fetches musics by Twit ID.
func (r *MusicsRepo) GetByTwitID(twitID string) ([]model.Music, error) {
	query := "SELECT id, twit_id, mp3 FROM musics WHERE twit_id = $1"
	rows, err := r.Db.Query(query, twitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musics []model.Music
	for rows.Next() {
		var music model.Music
		if err := rows.Scan(&music.ID, &music.TwitID, &music.MP3); err != nil {
			return nil, err
		}
		musics = append(musics, music)
	}
	return musics, nil
}

// GetByID fetches a music by its ID.
func (r *MusicsRepo) GetByID(id string) (*model.Music, error) {
	query := "SELECT id, twit_id, mp3 FROM musics WHERE id = $1"
	var music model.Music
	err := r.Db.QueryRow(query, id).Scan(&music.ID, &music.TwitID, &music.MP3)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &music, nil
}

// UpdateByID updates a music by its ID.
func (r *MusicsRepo) UpdateByID(id string, req *model.CreateMusicRequest) error {
	query := "UPDATE musics SET mp3 = $1 WHERE id = $2"
	_, err := r.Db.Exec(query, req.MP3, id)
	return err
}
