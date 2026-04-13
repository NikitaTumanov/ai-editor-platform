package grpcauth

import (
	"context"

	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
	"go.uber.org/zap"
)

type StorageRepository interface {
	CreateUser(ctx context.Context, in *storagepb.CreateUserRequest) (*storagepb.CreateUserResponse, error)
	FindUserByUsername(ctx context.Context, in *storagepb.FindUserByUsernameRequest) (*storagepb.FindUserByUsernameResponse, error)
	Close() error
}

type AuthHandler struct {
	logger *zap.Logger
	authpb.UnimplementedAuthServer
	storageRepo StorageRepository
}

func NewAuthHandler(logger *zap.Logger, storageRepo StorageRepository) *AuthHandler {
	return &AuthHandler{
		logger:      logger,
		storageRepo: storageRepo,
	}
}

func (s *AuthHandler) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.storageRepo.FindUserByUsername(ctx, &storagepb.FindUserByUsernameRequest{
		Login: in.Login,
	})

	return &authpb.LoginResponse{Token: "asacac"}, nil
}

func (s *AuthHandler) Register(ctx context.Context, in *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	s.storageRepo.CreateUser(ctx, &storagepb.CreateUserRequest{
		Login:    in.Login,
		Password: in.Password,
	})

	return &authpb.RegisterResponse{UserId: 1}, nil
}
