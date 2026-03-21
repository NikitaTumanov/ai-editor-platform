package service

import (
	"context"
	"errors"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/models"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
	"go.uber.org/zap"
)

type DocumentService struct {
	logger *zap.Logger
	repo   repository.DocumentRepository
}

func NewDocumentService(repo repository.DocumentRepository, logger *zap.Logger) *DocumentService {
	return &DocumentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *DocumentService) Create(ctx context.Context, document *models.Document, userID int64) (int64, error) {
	id, err := s.repo.Create(ctx, document, userID)
	if err != nil {
		s.logger.Error(
			"failed to create document",
			zap.String("documentName", document.Name),
			zap.Error(err),
		)
		return 0, err
	}

	s.logger.Info(
		"document created successfully",
		zap.Int64("documentID", id),
		zap.String("documentName", document.Name),
	)
	return id, nil
}

func (s *DocumentService) FindByID(ctx context.Context, id int64) (*models.Document, error) {
	document, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrDocumentNotFound) {
			s.logger.Warn(
				"document not found",
				zap.Int64("documentID", id),
			)
			return nil, err
		}
		s.logger.Error(
			"failed to find document by id",
			zap.Int64("documentID", id),
			zap.Error(err),
		)
		return nil, err

	}

	s.logger.Info(
		"document found successfully",
		zap.Int64("documentID", id),
		zap.String("documentName", document.Name),
	)
	return document, nil
}

func (s *DocumentService) FindByUserID(ctx context.Context, userID int64) ([]*models.Document, error) {
	documents, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, customerrors.ErrDocumentNotFound) {
			s.logger.Warn(
				"document not found",
				zap.Int64("userID", userID))
			return nil, err
		}
		s.logger.Error(
			"failed to find document by user id",
			zap.Int64("userID", userID),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"documents found successfully",
		zap.Int64("userID", userID),
	)
	return documents, nil
}

func (s *DocumentService) UpdateByID(ctx context.Context, id int64, newDocument *models.Document) error {
	err := s.repo.UpdateByID(ctx, id, newDocument)
	if err != nil {
		if errors.Is(err, customerrors.ErrDocumentNotFound) {
			s.logger.Warn(
				"document not found",
				zap.Int64("documentID", id),
			)
			return err
		}
		s.logger.Error(
			"failed to update document by id",
			zap.Int64("documentID", id),
			zap.Error(err),
		)
		return err
	}

	s.logger.Info(
		"document updated successfully",
		zap.Int64("documentID", id),
	)
	return nil
}

func (s *DocumentService) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrDocumentNotFound) {
			s.logger.Warn(
				"document not found",
				zap.Int64("documentID", id),
			)
			return err
		}
		s.logger.Error(
			"failed to delete document by id",
			zap.Int64("documentID", id),
			zap.Error(err),
		)
		return err
	}

	s.logger.Info(
		"document deleted successfully",
		zap.Int64("documentID", id),
	)
	return nil
}
