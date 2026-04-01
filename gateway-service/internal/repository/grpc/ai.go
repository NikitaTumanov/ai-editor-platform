package grpc

import (
	"context"
	"log"

	aipb "github.com/NikitaTumanov/ai-editor-platform/protos/ai_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AiServiceGrpc struct {
	Conn   *grpc.ClientConn
	Client aipb.AIClient
}

func NewAiServiceGrpc() *AiServiceGrpc {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8060", opts...)
	if err != nil {
		log.Fatalf("could not create grpc connection: %v", err)
	}

	client := aipb.NewAIClient(conn)

	return &AiServiceGrpc{
		Conn:   conn,
		Client: client,
	}
}

func (s *AiServiceGrpc) Question(ctx context.Context, in *aipb.QuestionRequest) (*aipb.QuestionResponse, error) {
	resp, err := s.Client.Question(ctx, in)
	return resp, err
}

func (s *AiServiceGrpc) UpdateDocumentById(ctx context.Context, in *aipb.UpdateDocumentByIdRequest) (*aipb.UpdateDocumentByIdResponse, error) {
	resp, err := s.Client.UpdateDocumentById(ctx, in)
	return resp, err
}

func (s *AiServiceGrpc) Close() error {
	return s.Conn.Close()
}
