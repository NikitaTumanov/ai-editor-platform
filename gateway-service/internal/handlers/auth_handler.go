package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/models"
	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthRepository interface {
	Register(ctx context.Context, in *authpb.RegisterRequest) (*authpb.RegisterResponse, error)
	Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error)
	ValidateToken(ctx context.Context, in *authpb.TokenRequest) (*authpb.TokenResponse, error)
}

type AuthHandler struct {
	logger   *zap.Logger
	authRepo AuthRepository
}

func NewAuthHandler(logger *zap.Logger, authRepo AuthRepository) *AuthHandler {
	return &AuthHandler{
		logger:   logger,
		authRepo: authRepo,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerReq models.RegisterDTO
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		h.logger.Warn(customerrors.ErrIncorrectJSON.Error(), zap.Error(err))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  customerrors.ErrIncorrectJSON.Error(),
			"fields": customerrors.FormatValidationError(err),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authRepo.Register(ctx, &authpb.RegisterRequest{
		Login:    registerReq.Login,
		Password: registerReq.Password,
	})
	if err != nil {
		h.logger.Error(
			"failed to register user",
			zap.String("login", registerReq.Login),
			zap.Error(err))

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": customerrors.ErrRegister.Error()})
		return
	}

	h.logger.Info(fmt.Sprintf("user %s registered with id: %d", registerReq.Login, resp.UserId))
	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq models.LoginDTO
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		h.logger.Warn(customerrors.ErrIncorrectJSON.Error(), zap.Error(err))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  customerrors.ErrIncorrectJSON.Error(),
			"fields": customerrors.FormatValidationError(err),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authRepo.Login(ctx, &authpb.LoginRequest{
		Login:    loginReq.Login,
		Password: loginReq.Password,
	})
	if err != nil {
		h.logger.Error(
			"failed to login",
			zap.String("login", loginReq.Login),
			zap.Error(err))

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": customerrors.ErrLogin.Error()})
		return
	}

	h.logger.Info(fmt.Sprintf("%s login successfully", loginReq.Login))
	c.JSON(http.StatusOK, resp)
}
