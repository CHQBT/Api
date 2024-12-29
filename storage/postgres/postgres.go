package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"milliy/config"
	"milliy/storage"
)

type postgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) storage.IStorage {
	return &postgresStorage{
		db: db,
	}
}

func ConnectionDb() (*sql.DB, error) {
	conf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Postgres.DB_HOST, conf.Postgres.DB_PORT, conf.Postgres.DB_USER, conf.Postgres.DB_NAME, conf.Postgres.DB_PASSWORD)
	db, err := sql.Open("postgres", conDb)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (p *postgresStorage) Close() {
	p.db.Close()
}

func (p *postgresStorage) Twit() storage.TwitStorage {
	return NewTwitRepo(p.db)
}

func (p *postgresStorage) Location() storage.LocationStorage {
	return NewLocationsRepo(p.db)
}

func (p *postgresStorage) Music() storage.MusicStorage {
	return NewMusicsRepo(p.db)
}

func (p *postgresStorage) Photo() storage.PhotoStorage {
	return NewPhotosRepo(p.db)
}

func (p *postgresStorage) Url() storage.UrlStorage {
	return NewURLsRepo(p.db)
}
func (p *postgresStorage) Video() storage.VideoStorage {
	return NewVideosRepo(p.db)
}
