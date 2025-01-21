package handler

import (
	"log/slog"
	"milliy/api/handler/twit"
	"milliy/api/handler/user"
	"milliy/api/middleware"
	"milliy/service"
	"milliy/upload"

	"github.com/casbin/casbin/v2"
)

type HandlerInterface interface {
	UserMethods() user.NewUser
	TwitMethods() twit.NewTwit
	EnforcerMethods() middleware.CasbinPermission
}

type Handler struct {
	User     *service.UserService
	Twit     *service.TwitService
	Log      *slog.Logger
	Enforcer *casbin.Enforcer
	MINIO    *upload.MinioUploader
}

func (h *Handler) EnforcerMethods() middleware.CasbinPermission {
	return middleware.NewCasbinPermission(h.Enforcer)
}

func (h *Handler) UserMethods() user.NewUser {
	return user.NewUsersMethods(h.User, h.Log)
}

func (h *Handler) TwitMethods() twit.NewTwit {
	return twit.NewTwitsMethods(h.User, h.Twit, h.Log, h.MINIO)
}
