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

	//router := gin.Default()
	//srv := &http.Server{
	//	Addr:         ":8040",
	//	Handler:      router,
	//	ReadTimeout:  10 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server forced to shutdown:", err)
	//}
	//log.Println("Server exiting")
}
