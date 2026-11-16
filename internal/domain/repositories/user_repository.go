package repositories

import (
	"context"

	"github.com/nujabuz/gocourse/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	UpdateUsername(ctx context.Context, id int64, newUsername string) error
	UpdatePassword(ctx context.Context, id int64, newPassword string) error
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) error
	GetByUsername(ctx context.Context, username string) error
}
