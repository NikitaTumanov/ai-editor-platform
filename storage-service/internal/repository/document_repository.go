package repository

import "github.com/jackc/pgx/v5/pgxpool"

type DocumentRepository interface {
}

type documentRepoPGX struct {
	pool *pgxpool.Pool
}

func NewDocumentRepoPGX(pool *pgxpool.Pool) DocumentRepository {
	return &documentRepoPGX{pool: pool}
}
