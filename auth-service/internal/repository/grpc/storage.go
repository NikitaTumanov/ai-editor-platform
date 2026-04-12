package grpc

import (
	"context"
	"log"

	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StorageServiceGrpc struct {
	Conn   *grpc.ClientConn
	Client storagepb.StorageClient
}

func NewStorageServiceGrpc() *StorageServiceGrpc {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8040", opts...)
	if err != nil {
		log.Fatalf("could not create grpc connection: %v", err)
	}

	client := storagepb.NewStorageClient(conn)

	return &StorageServiceGrpc{
		Conn:   conn,
		Client: client,
	}
}

func (s *StorageServiceGrpc) CreateUser(ctx context.Context, in *storagepb.CreateUserRequest) (*storagepb.CreateUserResponse, error) {
	resp, err := s.Client.CreateUser(ctx, in)
	return resp, err
}

func (s *StorageServiceGrpc) FindUserByUsername(ctx context.Context, in *storagepb.FindUserByUsernameRequest) (*storagepb.FindUserByUsernameResponse, error) {
	resp, err := s.Client.FindUserByUsername(ctx, in)
	return resp, err
}

func (s *StorageServiceGrpc) Close() error {
	return s.Conn.Close()
}
