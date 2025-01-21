package service

import (
	"context"
	"fmt"
	"log/slog"
	"milliy/logs"
	"milliy/model"
	"milliy/storage"
)

type UserService struct {
	storage storage.IStorage
	log     *slog.Logger
}

func NewUserService(storage storage.IStorage) *UserService {
	return &UserService{
		storage: storage,
		log:     logs.NewLogger(),
	}
}

func (u *UserService) Login(ctx context.Context, req *model.UserLogin) (*model.User, error) {
	u.log.Info("Login rpc started")
	resp, err := u.storage.User().GetUserByLogin(req.Login)
	if err != nil {
		u.log.Error(fmt.Sprintf("Error getting user: %v", err))
		return nil, err
	}
	b, err := u.storage.User().CheckPassword(req.Login, req.PasswordHash)
	if err != nil {
		u.log.Error(fmt.Sprintf("Error checking password: %v", err))
		return nil, err
	}
	if b != true {
		u.log.Error(fmt.Sprintf("password check failed: %v", b))
		return nil, fmt.Errorf("password hash mismatch")
	}
	u.log.Info("Login rpc finished")
	return resp, nil
}
