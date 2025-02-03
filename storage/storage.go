package storage

import "milliy/model"

type IStorage interface {
	Twit() TwitStorage
	Location() LocationStorage
	Music() MusicStorage
	Photo() PhotoStorage
	Url() UrlStorage
	Video() VideoStorage
	User() UserStorage
	Close()
}

type TwitStorage interface {
	CreateTwit(*model.CreateTwitRequest) (string, error)
	GetTwitByID(string) (*model.Twit, error)
	DeleteTwit(string) error
	AddReadersCount(string) error
	GetAllTwits() ([]string, error)
	GetTwitsByType(string) ([]string, error)
	GetMostViewedTwit(int) ([]string, error)
	GetLatestTwits(int) ([]string, error)
	SearchTwit(string) ([]string, error)
	GetUniqueTypes() ([]string, error)
}

type LocationStorage interface {
	Create(req *model.CreateLocationRequest) (string, error)
	DeleteByID(id string) error
	DeleteByTwitID(twitID string) error
	GetByTwitID(twitID string) ([]model.Location, error)
	GetByID(id string) (*model.Location, error)
	UpdateByID(id string, req *model.CreateLocationRequest) error
}

type MusicStorage interface {
	Create(req *model.CreateMusicRequest) (string, error)
	DeleteByID(id string) error
	DeleteByTwitID(twitID string) error
	GetByTwitID(twitID string) ([]model.Music, error)
	GetByID(id string) (*model.Music, error)
	UpdateByID(id string, req *model.CreateMusicRequest) error
}

type PhotoStorage interface {
	Create(req *model.CreatePhotoRequest) (string, error)
	DeleteByID(id string) error
	DeleteByTwitID(twitID string) error
	GetByTwitID(twitID string) ([]model.Photo, error)
	GetByID(id string) (*model.Photo, error)
	UpdateByID(id string, req *model.CreatePhotoRequest) error
}

type UrlStorage interface {
	Create(req *model.CreateURLRequest) (string, error)
	DeleteByID(id string) error
	DeleteByTwitID(twitID string) error
	GetByTwitID(twitID string) ([]model.URL, error)
	GetByID(id string) (*model.URL, error)
	UpdateByID(id string, req *model.CreateURLRequest) error
}

type VideoStorage interface {
	Create(req *model.CreateVideoRequest) (string, error)
	DeleteByID(id string) error
	DeleteByTwitID(twitID string) error
	GetByTwitID(twitID string) ([]model.Video, error)
	GetByID(id string) (*model.Video, error)
	UpdateByID(id string, req *model.CreateVideoRequest) error
}

type UserStorage interface {
	CheckPassword(login, password string) (bool, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByLogin(login string) (*model.User, error)
}
