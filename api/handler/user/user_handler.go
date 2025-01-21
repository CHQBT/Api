package user

import (
	"log/slog"
	"milliy/service"

	"github.com/gin-gonic/gin"
)

type newUsers struct {
	User *service.UserService
	Log  *slog.Logger
}

func NewUsersMethods(
	User *service.UserService,
	log *slog.Logger) NewUser {
	return &newUsers{
		User: User,
		Log:  log,
	}
}

type NewUser interface {
	Login(*gin.Context)
}
