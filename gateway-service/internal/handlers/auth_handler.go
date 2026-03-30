package handlers

import (
	"context"

	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthRepository interface {
	Register(ctx context.Context, in *authpb.RegisterRequest) (*authpb.RegisterResponse, error)
	Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error)
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
	h.authRepo.Register()
}

func (h *AuthHandler) Login(c *gin.Context) {
	h.authRepo.Login()
}
