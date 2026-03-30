package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DocumentRepository interface {
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

func (h *DocumentHandler) Add(c *gin.Context) {
	h.documentRepo.Add
}

func (h *DocumentHandler) Update(c *gin.Context) {
	h.documentRepo.Update
}
