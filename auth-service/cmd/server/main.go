package main

import (
	"context"
	"log"
	"net"

	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth"
	"google.golang.org/grpc"
)

type authServer struct {
	authpb.UnimplementedAuthServer
}

func (s *authServer) Login(_ context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{Token: "asacac"}, nil
}

func (s *authServer) Register(_ context.Context, in *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return &authpb.RegisterResponse{UserId: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8030")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	authpb.RegisterAuthServer(grpcServer, &authServer{})
	grpcServer.Serve(lis)
}
