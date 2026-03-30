package main

import (
	"context"
	"errors"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/handlers"
	zaplogger "github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/logger"
	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/repository/grpc"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	logger := zaplogger.New()
	defer logger.Sync()

	router := gin.Default()

	authRepo := grpc.NewAuthServiceGrpc()
	defer authRepo.Close()
	authHandler := handlers.NewAuthHandler(logger, authRepo)
	authAPI := router.Group("/auth")
	{
		authAPI.POST("/login", authHandler.Login)
		authAPI.POST("/register", authHandler.Register)
	}

	//aiRepo := grpc.NewAiServiceGrpc()
	//defer aiRepo.Close()
	//aiHandler := handlers.NewAiHandler(logger, aiRepo)
	aiAPI := router.Group("/ai")
	{
		aiAPI.GET("/ask", aiHandler.Ask)          // TODO
		aiAPI.PATCH("/updateFile", aiHandler.UpdateFile) // TODO
	}

	//documentRepo := grpc.NewDocumentServiceGrpc()
	//defer documentRepo.Close()
	//documentHandler := handlers.NewDocumentHandler(logger, documentRepo)
	documentAPI := router.Group("/document")
	{
		documentAPI.POST("/add", documentHandler.)     // TODO
		documentAPI.PATCH("/update", documentHandler.) // TODO
	}

	storageRepo := grpc.NewStorageServiceGrpc()
	defer storageRepo.Close()
	storageHandler := handlers.NewStorageHandler(logger, storageRepo)
	storageAPI := router.Group("/storage")
	{
		storageAPI.GET("/file/:id", storageHandler.DocumentByID)
		storageAPI.GET("/fileByUser/:user_id", storageHandler.DocumentsByUserID)
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("listen error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", zap.Error(err))
	}
	logger.Info("Server exiting")
}
