package service

import (
	"context"

	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/models"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UserByName(ctx context.Context, name string) (*models.User, error) {
	return s.repo.FindByUsername(ctx, name)
}
