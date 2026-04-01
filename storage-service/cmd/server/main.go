package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/database"
	zaplogger "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	gin.ForceConsoleColor()

	logger := zaplogger.New()
	defer logger.Sync()

	pool, err := database.NewPool("postgres://docsdbuser:docsdbpass@localhost:5432/docsdb")
	if err != nil {
		logger.Fatal("can't initialize database pool", zap.Error(err))
	}
	defer pool.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	log.Println("Server exiting")
}
