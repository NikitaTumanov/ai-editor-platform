package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/database"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	srv := &http.Server{
		Addr:         ":8090",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	pool, err := database.NewPool("postgres://docsdbuser:docsdbpass@localhost:5432/docsdb")
	if err != nil {
	}
	defer pool.Close()

	userRepo := repository.NewUserRepoPGX(pool)
	userService := service.NewUserService(userRepo)
	userService.UserByName(context.Background(), "user.Username()")

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
