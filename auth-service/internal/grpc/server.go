package grpcauth

import (
	"context"

	authpb "github.com/NikitaTumanov/ai-editor-platform/protos/auth_service"
)

type AuthServer struct {
	authpb.UnimplementedAuthServer
	//auth_service Auth
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

//type Auth interface {
//	Login(ctx context.Context, login string, password string) (token string, err error)
//	Register(ctx context.Context, login string, password string) (userID int64, err error)
//}

//func Register(gRPCServer *grpc.Server, auth_service Auth) {
//	authpb.RegisterAuthServer(gRPCServer, &authServer{auth_service: auth_service})
//}

func (s *AuthServer) Login(_ context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{Token: "asacac"}, nil
}

func (s *AuthServer) Register(_ context.Context, in *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return &authpb.RegisterResponse{UserId: 1}, nil
}
