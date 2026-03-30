package grpc

import (
	"context"
	"log"

	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceGrpc struct {
	Conn   *grpc.ClientConn
	Client authpb.AuthClient
}

func NewAuthServiceGrpc() *AuthServiceGrpc {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8070", opts...)
	if err != nil {
		log.Fatalf("could not create grpc connection: %v", err)
	}

	client := authpb.NewAuthClient(conn)

	return &AuthServiceGrpc{
		Conn:   conn,
		Client: client,
	}
}

func (s *AuthServiceGrpc) Register(ctx context.Context, in *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	resp, err := s.Client.Register(ctx, in)
	return resp, err

}

func (s *AuthServiceGrpc) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	resp, err := s.Client.Login(ctx, in)
	return resp, err
}

func (s *AuthServiceGrpc) Close() error {
	return s.Conn.Close()
}
