package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"milliy/model"
	"milliy/storage"
	"time"
)

type TwitRepo struct {
	Db *sql.DB
}

func NewTwitRepo(db *sql.DB) storage.TwitStorage {
	return &TwitRepo{Db: db}
}

func (r *TwitRepo) CreateTwit(req *model.CreateTwitRequest) (string, error) {
	query := `
		INSERT INTO twit (
			user_id, 
			publisher_FIO,
			type,
			title, 
			texts, 
			readers_count
		) VALUES ($1, $2, $3, $4, $5, 0) 
		RETURNING id`

	var id string
	err := r.Db.QueryRow(
		query,
		req.UserID,
		req.PublisherFIO,
		req.Type,
		req.Title,
		req.Texts,
	).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetTwitByID fetches a twit by its ID.
func (r *TwitRepo) GetTwitByID(id string) (*model.Twit, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			publisher_FIO,
			type,
			title, 
			texts, 
			readers_count,
			created_at
		FROM twit 
		WHERE id = $1 AND deleted_at = 0`

	var twit model.Twit
	err := r.Db.QueryRow(query, id).Scan(
		&twit.ID,
		&twit.UserID,
		&twit.PublisherFIO,
		&twit.Type,
		&twit.Title,
		&twit.Texts,
		&twit.ReadersCount,
		&twit.CreatedAt,
	)
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

func (p *TwitRepo) GetAllTwits() ([]string, error) {
	var ids []string
	query := `SELECT id FROM twit WHERE deleted_at = 0`

	rows, err := p.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (p *TwitRepo) GetTwitsByType(twitType string) ([]string, error) {
	var ids []string
	query := `SELECT id FROM twit WHERE type = $1 AND deleted_at = 0`

	rows, err := p.Db.Query(query, twitType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (p *TwitRepo) GetMostViewedTwit(limit int) ([]string, error) {
	var ids []string
	query := `
		SELECT id
		FROM twit 
		WHERE deleted_at = 0
		ORDER BY readers_count DESC 
		LIMIT $1`

	rows, err := p.Db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (p *TwitRepo) GetLatestTwits(limit int) ([]string, error) {
	var ids []string
	query := `
		SELECT id
		FROM twit 
		WHERE deleted_at = 0
		ORDER BY created_at DESC 
		LIMIT $1`

	rows, err := p.Db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (p *TwitRepo) SearchTwit(keyword string) ([]string, error) {
	var ids []string
	query := `
		SELECT id
		FROM twit 
		WHERE deleted_at = 0 AND
			(title ILIKE '%' || $1 || '%' OR 
			 texts ILIKE '%' || $1 || '%' OR
			 publisher_FIO ILIKE '%' || $1 || '%')
	`

	rows, err := p.Db.Query(query, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (p *TwitRepo) GetUniqueTypes() ([]string, error) {
	var types []string
	query := `
        SELECT DISTINCT type 
        FROM twit 
        WHERE deleted_at = 0
        ORDER BY type` // Optional sorting for consistent order

	rows, err := p.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}

func (p *TwitRepo) AddMainTwit(twitID, start_time, end_time string) error {
	// Kutilayotgan vaqt formati: "dd-mm-yyyy hh-mm"
	layout := "02-01-2006 15-04" // Go-da 2006-01-02 15:04:05 asosiy format

	// start_time va end_time ni `time.Time` ga o‘giramiz
	startTimeParsed, err := time.Parse(layout, start_time)
	if err != nil {
		return fmt.Errorf("invalid start_time format: %w", err)
	}

	endTimeParsed, err := time.Parse(layout, end_time)
	if err != nil {
		return fmt.Errorf("invalid end_time format: %w", err)
	}

	query := `
        INSERT INTO saved (twit_id, start_time, end_time, created_at, updated_at, deleted_at)
        VALUES ($1, $2, $3, NOW(), NOW(), 0)
    `

	_, err = p.Db.Exec(query, twitID, startTimeParsed, endTimeParsed)
	if err != nil {
		return fmt.Errorf("failed to insert main twit: %w", err)
	}

	return nil
}

func (p *TwitRepo) GetMainTwit() ([]string, error) {
	query := `
        SELECT twit_id 
        FROM saved 
        WHERE deleted_at = 0 
          AND start_time <= NOW() 
          AND end_time >= NOW()
    `

	rows, err := p.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get main twit: %w", err)
	}
	defer rows.Close()

	var twitIDs []string
	for rows.Next() {
		var twitID string
		if err := rows.Scan(&twitID); err != nil {
			return nil, err
		}
		twitIDs = append(twitIDs, twitID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return twitIDs, nil
}

func (p *TwitRepo) DeleteMainTwit(twitID string) error {
	query := `
        UPDATE saved
        SET deleted_at = EXTRACT(EPOCH FROM NOW())
        WHERE twit_id = $1 AND deleted_at = 0
    `

	result, err := p.Db.Exec(query, twitID)
	if err != nil {
		return fmt.Errorf("failed to delete main twit: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("main twit not found or already deleted")
	}

	return nil
}
