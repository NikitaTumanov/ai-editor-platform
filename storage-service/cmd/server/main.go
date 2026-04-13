package main

import (
	"log"
	"net"

	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/database"
	storagegrpc "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/grpc"
	zaplogger "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/logger"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := zaplogger.New()
	defer logger.Sync()

	pool, err := database.NewPool("postgres://docsdbuser:docsdbpass@localhost:5432/docsdb")
	if err != nil {
		logger.Fatal("can't initialize database pool", zap.Error(err))
	}
	defer pool.Close()

	userRepo := repository.NewUserRepoPGX(pool)
	documentRepo := repository.NewDocumentRepoPGX(pool)

	lis, err := net.Listen("tcp", ":8040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	storagepb.RegisterStorageServer(grpcServer, storagegrpc.NewStorageHandler(logger, userRepo, documentRepo))
	grpcServer.Serve(lis)
}
