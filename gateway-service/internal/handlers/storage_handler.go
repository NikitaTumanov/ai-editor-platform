package handlers

import (
	"context"

	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StorageRepository interface {
	DocumentByID(ctx context.Context, in *storagepb.GetDocumentByIdRequest) (*storagepb.GetDocumentByIdResponse, error)
	DocumentsByUserID(ctx context.Context, in *storagepb.GetDocumentsByUserIdRequest) (*storagepb.GetDocumentsByUserIdResponse, error)
}

type StorageHandler struct {
	logger      *zap.Logger
	storageRepo StorageRepository
}

func NewStorageHandler(logger *zap.Logger, storageRepo StorageRepository) *StorageHandler {
	return &StorageHandler{
		logger:      logger,
		storageRepo: storageRepo,
	}
}

func (h *StorageHandler) DocumentByID(c *gin.Context) {
	//h.storageRepo.DocumentByID()
}

func (h *StorageHandler) DocumentsByUserID(c *gin.Context) {
	//h.storageRepo.DocumentsByUserID()
}
