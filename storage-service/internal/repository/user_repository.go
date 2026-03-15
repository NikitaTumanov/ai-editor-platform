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

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
	UpdateByUsername(ctx context.Context, username string, newUser *models.User) error
	Delete(ctx context.Context, username string) error
}

type userRepoPGX struct {
	pool *pgxpool.Pool
}

func NewUserRepoPGX(pool *pgxpool.Pool) UserRepository {
	return &userRepoPGX{pool: pool}
}

func (r *userRepoPGX) Create(ctx context.Context, user *models.User) (int64, error) {
	var id int64
	err := r.pool.QueryRow(ctx, "INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id;",
		user.Login,
		user.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user %s into database: %s", user.Login, err)
	}

	return id, nil
}

func (r *userRepoPGX) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var (
		id              int64
		login, password string
	)
	err := r.pool.QueryRow(ctx, "SELECT id, login, password_hash FROM users WHERE login=$1;",
		username).Scan(&id, &login, &password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, customerrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user by username %s: %s", username, err)
	}

	return &models.User{ID: id, Login: login, Password: password}, nil
}

func (r *userRepoPGX) FindByID(ctx context.Context, id int64) (*models.User, error) {
	var login, password string
	err := r.pool.QueryRow(ctx, "SELECT id, login, password_hash FROM users WHERE id=$1;",
		id).Scan(&id, &login, &password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, customerrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user by id %d: %s", id, err)
	}

	return &models.User{ID: id, Login: login, Password: password}, nil
}

func (r *userRepoPGX) UpdateByUsername(ctx context.Context, username string, newUser *models.User) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction on update user by username %s: %s", username, err)
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	commandTag, err := tx.Exec(ctx, "UPDATE users SET login=$1, password_hash=$2 WHERE login=$3;",
		newUser.Login,
		newUser.Password,
		username)
	if err != nil {
		return fmt.Errorf("failed to update user by username %s: %s", username, err)
	}

	rows := commandTag.RowsAffected()
	if rows == 0 {
		return customerrors.ErrUserNotFound
	}
	if rows > 1 {
		return fmt.Errorf("failed to update user by username %s: too many rows affected", username)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction on update user by username %s: %s", username, err)
	}
	return nil
}

func (r *userRepoPGX) Delete(ctx context.Context, username string) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction on delete user by username %s: %s", username, err)
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	commandTag, err := tx.Exec(ctx, "DELETE FROM users WHERE login=$1;", username)
	if err != nil {
		return fmt.Errorf("failed to delete user by username %s: %s", username, err)
	}

	rows := commandTag.RowsAffected()
	if rows == 0 {
		return customerrors.ErrUserNotFound
	}
	if rows > 1 {
		return fmt.Errorf("failed to delete user by username %s: too many rows affected", username)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction on delete user by username %s: %s", username, err)
	}
	return nil
}
