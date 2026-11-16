// Package repositories
package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nujabuz/gocourse/internal/domain/entities"
	"github.com/nujabuz/gocourse/internal/domain/repositories"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) repositories.UserRepository {
	userRepository := new(userRepository)
	userRepository.db = dbPool

	return userRepository
}

func (repo *userRepository) Create(ctx context.Context, user *entities.User) error {
	query := "INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4);"
	_, err := repo.db.Exec(ctx, query, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)

	return err
}

func (repo *userRepository) UpdateUsername(ctx context.Context, id int64, newUsername string) error {
	query := "UPDATE users SET username = $2 WHERE id = $1"
	_, err := repo.db.Exec(ctx, query, id, newUsername)

	return err
}

func (repo *userRepository) UpdatePassword(ctx context.Context, id int64, newPassword string) error {
	query := "UPDATE users SET password = $2 WHERE id = $1"
	_, err := repo.db.Exec(ctx, query, id, newPassword)

	return err
}

func (repo *userRepository) DeleteByID(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := repo.db.Exec(ctx, query, id)

	return err
}

func (repo *userRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	query := "SELECT id, username, password, created_at, updated_at FROM users WHERE id = $1"
	row := repo.db.QueryRow(ctx, query, id)

	user := new(entities.User)
	if err := row.Scan(user.ID, user.Username, user.Password, user.CreatedAt, user.UpdatedAt); err != nil {
		return nil, nil
	}

	return user, nil
}

func (repo *userRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := "SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1"
	row := repo.db.QueryRow(ctx, query, username)

	user := new(entities.User)
	if err := row.Scan(user.ID, user.Username, user.Password, user.CreatedAt, user.UpdatedAt); err != nil {
		return nil, nil
	}

	return user, nil
}
