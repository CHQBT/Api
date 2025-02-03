package twit

import (
	"log/slog"
	"milliy/service"
	"milliy/upload"

	"github.com/gin-gonic/gin"
)

type newTwits struct {
	User  *service.UserService
	Twit  *service.TwitService
	Log   *slog.Logger
	MINIO *upload.MinioUploader
}

func NewTwitsMethods(
	User *service.UserService,
	Twit *service.TwitService,
	log *slog.Logger,
	MINIO *upload.MinioUploader) NewTwit {
	return &newTwits{
		User:  User,
		Twit:  Twit,
		Log:   log,
		MINIO: MINIO,
	}
}

type NewTwit interface {
	CreateTwit(c *gin.Context)
	GetTwit(c *gin.Context)
	DeleteTwit(*gin.Context)
	AddReadersCount(*gin.Context)
	GetAllTwits(*gin.Context)
	GetTwitsByType(*gin.Context)
	GetMostViewedTwits(*gin.Context)
	GetLatestTwits(*gin.Context)
	SearchTwit(*gin.Context)
	CreateLocation(c *gin.Context)
	CreateUrl(c *gin.Context)
	CreatePhoto(c *gin.Context)
	CreateVideo(c *gin.Context)
	CreateMusic(c *gin.Context)
	GetUniqueTypes(c *gin.Context)
}
