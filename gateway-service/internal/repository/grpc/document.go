package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	documentpb "github.com/NikitaTumanov/ai-editor-platform/protos/document_service"
)

type DocumentServiceGrpc struct {
	Conn   *grpc.ClientConn
	Client documentpb.DocumentClient
}

func NewDocumentServiceGrpc() *DocumentServiceGrpc {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8050", opts...)
	if err != nil {
		log.Fatalf("could not create grpc connection: %v", err)
	}

	client := documentpb.NewDocumentClient(conn)

	return &DocumentServiceGrpc{
		Conn:   conn,
		Client: client,
	}
}

func (s *DocumentServiceGrpc) AddDocument(ctx context.Context, in *documentpb.AddDocumentRequest) (*documentpb.AddDocumentResponse, error) {
	resp, err := s.Client.AddDocument(ctx, in)
	return resp, err
}

func (s *DocumentServiceGrpc) UpdateDocumentById(ctx context.Context, in *documentpb.UpdateDocumentByIdRequest) (*documentpb.UpdateDocumentByIdResponse, error) {
	resp, err := s.Client.UpdateDocumentById(ctx, in)
	return resp, err
}

func (s *DocumentServiceGrpc) Close() error {
	return s.Conn.Close()
}
