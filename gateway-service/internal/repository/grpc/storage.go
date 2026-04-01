package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
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

func (s *StorageServiceGrpc) DocumentByID(ctx context.Context, in *storagepb.GetDocumentByIdRequest) (*storagepb.GetDocumentByIdResponse, error) {
	res, err := s.Client.GetDocumentById(ctx, in)
	return res, err
}

func (s *StorageServiceGrpc) DocumentsByUserID(ctx context.Context, in *storagepb.GetDocumentsByUserIdRequest) (*storagepb.GetDocumentsByUserIdResponse, error) {
	res, err := s.Client.GetDocumentsByUserId(ctx, in)
	return res, err
}

func (s *StorageServiceGrpc) Close() error {
	return s.Conn.Close()
}
