package service

import (
	"context"
	"errors"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/models"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
)

type UserService struct {
	//logger
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UserByName(ctx context.Context, name string) (*models.User, error) {
	user, err := s.repo.FindByUsername(ctx, name)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {

		}

	}

	return user, nil
}
