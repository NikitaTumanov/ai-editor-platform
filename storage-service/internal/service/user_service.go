package service

import (
	"context"
	"errors"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/models"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
	"go.uber.org/zap"
)

type UserService struct {
	logger *zap.Logger
	repo   repository.UserRepository
}

func NewUserService(repo repository.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.User) (int64, error) {
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		s.logger.Error(
			"failed to create user",
			zap.String("userLogin", user.Login),
			zap.Error(err),
		)
		return 0, err
	}

	s.logger.Info(
		"user created successfully",
		zap.Int64("userID", id),
		zap.String("userLogin", user.Login),
	)
	return id, nil
}

func (s *UserService) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if errors.Is(err, customerrors.ErrUserNotFound) {
		s.logger.Warn(
			"user not found",
			zap.String("userLogin", username),
		)
		return nil, err
	}
	if err != nil {
		s.logger.Error(
			"failed to find user by username",
			zap.String("userLogin", username),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"user found successfully",
		zap.Int64("userID", user.ID),
		zap.String("userLogin", user.Login),
	)
	return user, nil
}

func (s *UserService) FindByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {
			s.logger.Warn(
				"user not found",
				zap.Int64("userID", id),
			)
			return nil, err
		}
		s.logger.Error(
			"failed to find user by id",
			zap.Int64("userID", id),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"user found successfully",
		zap.Int64("userID", user.ID),
		zap.String("userLogin", user.Login),
	)
	return user, nil
}

func (s *UserService) UpdateByUsername(ctx context.Context, username string, newUser *models.User) error {
	err := s.repo.UpdateByUsername(ctx, username, newUser)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {
			s.logger.Warn(
				"user not found",
				zap.String("userLogin", username),
			)
			return err
		}
		s.logger.Error(
			"failed to update user by username",
			zap.String("userLogin", username),
			zap.Error(err),
		)
		return err
	}

	s.logger.Info(
		"user updated successfully",
		zap.String("userLogin", username),
	)
	return nil
}

func (s *UserService) Delete(ctx context.Context, username string) error {
	err := s.repo.Delete(ctx, username)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {
			s.logger.Warn(
				"user not found",
				zap.String("userLogin", username),
			)
			return err
		}
		s.logger.Error(
			"failed to delete user by username",
			zap.String("userLogin", username),
			zap.Error(err),
		)
		return err
	}

	s.logger.Info(
		"user deleted successfully",
		zap.String("userLogin", username),
	)
	return nil
}
