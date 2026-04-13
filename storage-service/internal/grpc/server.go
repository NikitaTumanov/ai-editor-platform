package storagegrpc

import (
	"context"

	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/models"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
	"go.uber.org/zap"
)

type StorageHandler struct {
	logger *zap.Logger
	storagepb.UnimplementedStorageServer
	userRepo     repository.UserRepository
	documentRepo repository.DocumentRepository
}

func NewStorageHandler(logger *zap.Logger, userRepo repository.UserRepository, documentRepo repository.DocumentRepository) *StorageHandler {
	return &StorageHandler{
		logger:       logger,
		userRepo:     userRepo,
		documentRepo: documentRepo,
	}
}

func (s *StorageHandler) CreateUser(ctx context.Context, in *storagepb.CreateUserRequest) (*storagepb.CreateUserResponse, error) {
	s.userRepo.Create(context.Background(), &models.User{Login: in.Login, Password: in.Password})

	return &storagepb.CreateUserResponse{UserId: 1}, nil
}

func (s *StorageHandler) FindUserByUsername(ctx context.Context, in *storagepb.FindUserByUsernameRequest) (*storagepb.FindUserByUsernameResponse, error) {
	s.userRepo.FindByUsername(context.Background(), in.Login)
	return &storagepb.FindUserByUsernameResponse{UserId: 1, Login: "ascasdc", Password: "acascas"}, nil
}
