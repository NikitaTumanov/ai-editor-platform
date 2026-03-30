package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AiRepository interface {
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

func (h *AiHandler) Ask(c *gin.Context) {

}

func (h *AiHandler) UpdateFile(c *gin.Context) {

}
