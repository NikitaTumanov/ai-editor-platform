package handlers

import (
	"context"

	documentpb "github.com/NikitaTumanov/ai-editor-platform/protos/document_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DocumentRepository interface {
	AddDocument(ctx context.Context, in *documentpb.AddDocumentRequest) (*documentpb.AddDocumentResponse, error)
	UpdateDocumentById(ctx context.Context, in *documentpb.UpdateDocumentByIdRequest) (*documentpb.UpdateDocumentByIdResponse, error)
}

type DocumentHandler struct {
	logger       *zap.Logger
	documentRepo DocumentRepository
}

func NewDocumentHandler(logger *zap.Logger, documentRepo DocumentRepository) *DocumentHandler {
	return &DocumentHandler{
		logger:       logger,
		documentRepo: documentRepo,
	}
}

func (h *DocumentHandler) AddDocument(c *gin.Context) {
	// TODO Вызвать проверку токена и получить ID пользователя
	//userID := int64(3)

}

func (h *DocumentHandler) UpdateDocumentById(c *gin.Context) {
	//h.documentRepo.UpdateDocumentById()
}
