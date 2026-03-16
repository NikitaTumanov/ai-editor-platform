package zaplogger

import (
	"log"

	"go.uber.org/zap"
)

func New() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	logger = logger.With(zap.String("service", "storage-server"))
	return logger
}
