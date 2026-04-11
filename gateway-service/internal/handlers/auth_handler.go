package handlers

import (
	"context"
	"fmt"
	"net/http"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/models"
	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func formatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range ve {
			field := fieldErr.Field()

			switch fieldErr.Tag() {
			case "required":
				errors[field] = "field is required"
			case "min":
				errors[field] = "too short"
			case "max":
				errors[field] = "too long"
			default:
				errors[field] = "invalid value"
			}
		}
	}

	return errors
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerReq models.RegisterDTO
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		h.logger.Warn("failed to parse body", zap.Error(err))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  customerrors.ErrIncorrectJSON.Error(),
			"fields": formatValidationError(err),
		})
		return
	}

	//if registerReq.Login == "" {
	//	h.logger.Warn("user login is empty")
	//
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrEmptyLogin.Error()})
	//	return
	//}
	//if len([]rune(registerReq.Login)) > settings.MaxLoginLen {
	//	h.logger.Warn(customerrors.ErrLongLogin.Error())
	//
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrLongLogin.Error()})
	//	return
	//}
	//if len([]rune(registerReq.Login)) < settings.MinLoginLen {
	//	h.logger.Warn(customerrors.ErrShortLogin.Error())
	//
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrShortLogin.Error()})
	//	return
	//}

	//if registerReq.Password == "" {
	//	h.logger.Warn("user password is empty")
	//
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrEmptyPassword.Error()})
	//	return
	//}
	//if len([]rune(registerReq.Password)) > settings.MaxPasswordLen {
	//	h.logger.Warn(customerrors.ErrLongPassword.Error())
	//
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrLongPassword.Error()})
	//	return
	//}
	//if len([]rune(registerReq.Password)) < settings.MinPasswordLen {
	//	h.logger.Warn(customerrors.ErrShortPassword.Error())
	//
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": customerrors.ErrShortPassword.Error()})
	//	return
	//}

	ctx := context.Background() // Сменить на нормальный
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
	ctx := context.Background()
	h.authRepo.Login(ctx, &authpb.LoginRequest{})
}
