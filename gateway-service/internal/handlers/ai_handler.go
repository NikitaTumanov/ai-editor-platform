package handlers

import (
	"context"

	aipb "github.com/NikitaTumanov/ai-editor-platform/protos/ai_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AiRepository interface {
	Question(ctx context.Context, in *aipb.QuestionRequest) (*aipb.QuestionResponse, error)
	UpdateDocumentById(ctx context.Context, in *aipb.UpdateDocumentByIdRequest) (*aipb.UpdateDocumentByIdResponse, error)
}

type AiHandler struct {
	logger *zap.Logger
	aiRepo AiRepository
}

func NewAiHandler(logger *zap.Logger, aiRepo AiRepository) *AiHandler {
	return &AiHandler{
		logger: logger,
		aiRepo: aiRepo,
	}
}

func (h *AiHandler) Question(c *gin.Context) {
	h.aiRepo.Question()
}

func (h *AiHandler) UpdateDocumentById(c *gin.Context) {
	h.aiRepo.UpdateDocumentById()
}
