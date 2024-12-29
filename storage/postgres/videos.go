package postgres

import (
	"database/sql"
	"errors"
	"milliy/model"
	"milliy/storage"
)

type VideosRepo struct {
	Db *sql.DB
}

func NewVideosRepo(db *sql.DB) storage.VideoStorage {
	return &VideosRepo{Db: db}
}

// Create inserts a new video record.
func (r *VideosRepo) Create(req *model.CreateVideoRequest) (string, error) {
	query := "INSERT INTO videos (twit_id, video) VALUES ($1, $2) RETURNING id"
	var id string
	err := r.Db.QueryRow(query, req.TwitID, req.Video).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// DeleteByID deletes a video by its ID.
func (r *VideosRepo) DeleteByID(id string) error {
	query := "DELETE FROM videos WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	return err
}

// DeleteByTwitID deletes videos associated with a specific Twit ID.
func (r *VideosRepo) DeleteByTwitID(twitID string) error {
	query := "DELETE FROM videos WHERE twit_id = $1"
	_, err := r.Db.Exec(query, twitID)
	return err
}

// GetByTwitID fetches videos by Twit ID.
func (r *VideosRepo) GetByTwitID(twitID string) ([]model.Video, error) {
	query := "SELECT id, twit_id, video FROM videos WHERE twit_id = $1"
	rows, err := r.Db.Query(query, twitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []model.Video
	for rows.Next() {
		var video model.Video
		if err := rows.Scan(&video.ID, &video.TwitID, &video.Video); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

// GetByID fetches a video by its ID.
func (r *VideosRepo) GetByID(id string) (*model.Video, error) {
	query := "SELECT id, twit_id, video FROM videos WHERE id = $1"
	var video model.Video
	err := r.Db.QueryRow(query, id).Scan(&video.ID, &video.TwitID, &video.Video)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &video, nil
}

// UpdateByID updates a video by its ID.
func (r *VideosRepo) UpdateByID(id string, req *model.CreateVideoRequest) error {
	query := "UPDATE videos SET video = $1 WHERE id = $2"
	_, err := r.Db.Exec(query, req.Video, id)
	return err
}
