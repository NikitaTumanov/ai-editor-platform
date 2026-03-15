package main

import (
	"log"
	"net"

	grpcauth "github.com/NikitaTumanov/ai-editor-platform/auth-service/internal/grpc"
	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8030")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	authpb.RegisterAuthServer(grpcServer, grpcauth.NewAuthServer())
	grpcServer.Serve(lis)
}
