package service

import (
	"context"
	"milliy/config"
	pb "milliy/generated/api"
	"milliy/storage"

	"milliy/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	storage storage.IStorage
	cfg     *config.Config
}

func NewUserService(storage storage.IStorage) *userService {
	return &userService{
		storage: storage,
		cfg:     config.Load(),
	}
}

func (s *userService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	// Check password
	ok, err := s.storage.User().CheckPassword(req.Login, req.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// Get user
	user, err := s.storage.User().GetUserByLogin(req.Login)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// Create response
	response := &pb.LoginRes{
		Id:   user.ID,
		Role: user.Role,
	}

	// Generate token
	tokenString, err := auth.GeneratedRefreshJWTToken(response)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	response.Token = tokenString
	return response, nil
}
