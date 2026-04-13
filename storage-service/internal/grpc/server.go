package storagegrpc

import (
	"context"

	storagepb "github.com/NikitaTumanov/ai-editor-platform/protos/storage_service"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type StorageHandler struct {
	logger *zap.Logger
	storagepb.UnimplementedStorageServer
	pool *pgxpool.Pool
}

func NewStorageHandler(logger *zap.Logger, pool *pgxpool.Pool) *StorageHandler {
	return &StorageHandler{
		logger: logger,
		pool:   pool,
	}
}

func (s *StorageHandler) CreateUser(ctx context.Context, in *storagepb.) {

}
