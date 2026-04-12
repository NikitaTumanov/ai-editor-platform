package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/models"
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
	// TODO Вызвать проверку токена и получить ID пользователя

	var questionReq models.QuestionDTO
	err := c.ShouldBindJSON(&questionReq)
	if err != nil {
		h.logger.Warn(customerrors.ErrIncorrectJSON.Error(), zap.Error(err))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  customerrors.ErrIncorrectJSON.Error(),
			"fields": customerrors.FormatValidationError(err),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := h.aiRepo.Question(ctx, &aipb.QuestionRequest{
		Question: questionReq.Question,
	})
	if err != nil {
		h.logger.Error(
			"failed to execute question",
			zap.Error(err),
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": customerrors.ErrQuestion.Error()})
		return
	}

	h.logger.Info("question validated successfully")
	c.JSON(http.StatusOK, resp)
}

func (h *AiHandler) UpdateDocumentById(c *gin.Context) {
	// TODO Вызвать проверку токена и получить ID пользователя
	userID := int64(3)

	var updateDocumentByIDReq models.UpdateDocumentByIDDTO
	err := c.ShouldBindJSON(&updateDocumentByIDReq)
	if err != nil {
		h.logger.Warn(customerrors.ErrIncorrectJSON.Error(), zap.Error(err))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  customerrors.ErrIncorrectJSON.Error(),
			"fields": customerrors.FormatValidationError(err),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resp, err := h.aiRepo.UpdateDocumentById(ctx, &aipb.UpdateDocumentByIdRequest{
		DocumentId: updateDocumentByIDReq.DocumentID,
		UserId:     userID,
		Promt:      updateDocumentByIDReq.Promt,
	})
	if err != nil {
		h.logger.Error(
			"failed to update document",
			zap.Error(err))

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": customerrors.ErrUpdateDocumentByID.Error()})
		return
	}

	h.logger.Info(fmt.Sprintf("document with id %d updated successfully", userID))
	c.JSON(http.StatusOK, resp)
}
