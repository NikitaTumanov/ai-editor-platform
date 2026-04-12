package main

import (
	"log"
	"net"

	authgrpc "github.com/NikitaTumanov/ai-editor-platform/auth-service/internal/handlers/grpc"
	zaplogger "github.com/NikitaTumanov/ai-editor-platform/auth-service/internal/logger"
	storagegrpc "github.com/NikitaTumanov/ai-editor-platform/auth-service/internal/repository/grpc"
	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
	"google.golang.org/grpc"
)

func main() {
	logger := zaplogger.New()
	defer logger.Sync()

	lis, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	storageRepo := storagegrpc.NewStorageServiceGrpc()
	defer storageRepo.Close()
	authpb.RegisterAuthServer(grpcServer, authgrpc.NewAuthHandler(logger, storageRepo))
	grpcServer.Serve(lis)
}
