package repository

import (
	"context"
	"errors"
	"fmt"

	customerrors "github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/errors"
	"github.com/NikitaTumanov/ai-editor-platform/storage-service/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *models.Document, userID int64) (int64, error)
	FindByID(ctx context.Context, id int64) (*models.Document, error)
	FindByUserID(ctx context.Context, userID int64) ([]*models.Document, error)
	UpdateByID(ctx context.Context, id int64, newDocument *models.Document) error
	Delete(ctx context.Context, id int64) error
}

type documentRepoPGX struct {
	pool *pgxpool.Pool
}

func NewDocumentRepoPGX(pool *pgxpool.Pool) DocumentRepository {
	return &documentRepoPGX{pool: pool}
}

func (r *documentRepoPGX) Create(ctx context.Context, document *models.Document, userID int64) (int64, error) {
	var id int64
	err := r.pool.QueryRow(ctx, "INSERT INTO documents (name, data, user_id) VALUES ($1, $2, $3) RETURNING id;",
		document.Name,
		document.Data,
		userID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert document %s into database: %s", document.Name, err)
	}

	return id, nil
}

func (r *documentRepoPGX) FindByID(ctx context.Context, id int64) (*models.Document, error) {
	var (
		name   string
		data   []byte
		userID int64
	)
	err := r.pool.QueryRow(ctx, "SELECT name, data, user_id FROM documents WHERE id=$1;",
		id).Scan(&name, &data, &userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, customerrors.ErrDocumentNotFound
		}
		return nil, fmt.Errorf("failed to find document by id %d: %s", id, err)
	}

	return &models.Document{Name: name, Data: data, UserID: userID}, nil
}

func (r *documentRepoPGX) FindByUserID(ctx context.Context, userID int64) ([]*models.Document, error) {
	var documents []*models.Document
	rows, err := r.pool.Query(ctx, "SELECT id, name, data FROM documents WHERE user_id=$1", userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, customerrors.ErrDocumentNotFound
		}
		return nil, fmt.Errorf("failed to find documents by user id %d: %s", userID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var document models.Document
		err := rows.Scan(&document.ID, &document.Name, &document.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to parse documents from rows by user id %d: %s", userID, err)
		}

		documents = append(documents, &document)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to find documents by user id %d: %s", userID, err)
	}

	return documents, nil
}

func (r *documentRepoPGX) UpdateByID(ctx context.Context, id int64, newDocument *models.Document) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction on update document by id %d: %s", id, err)
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	commandTag, err := tx.Exec(ctx, "UPDATE documents SET name=$1, data=$2, user_id=$3 WHERE id=$4;",
		newDocument.Name,
		newDocument.Data,
		newDocument.UserID,
		id)
	if err != nil {
		return fmt.Errorf("failed to update document by id %d: %s", id, err)
	}

	rows := commandTag.RowsAffected()
	if rows == 0 {
		return customerrors.ErrDocumentNotFound
	}
	if rows > 1 {
		return fmt.Errorf("failed to update document by id %d: too many rows affected", id)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction on update document by id %d: %s", id, err)
	}
	return nil
}

func (r *documentRepoPGX) Delete(ctx context.Context, id int64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction on delete document by id %d: %s", id, err)
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	commandTag, err := tx.Exec(ctx, "DELETE FROM documents WHERE id=$1;", id)
	if err != nil {
		return fmt.Errorf("failed to delete document by id %d: %s", id, err)
	}

	rows := commandTag.RowsAffected()
	if rows == 0 {
		return customerrors.ErrDocumentNotFound
	}
	if rows > 1 {
		return fmt.Errorf("failed to delete document by id %d: %s", id, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction on delete document by id %d: %s", id, err)
	}
	return nil
}
