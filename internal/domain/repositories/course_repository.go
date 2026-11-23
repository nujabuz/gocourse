package repositories

import (
	"context"

	"github.com/nujabuz/gocourse/internal/domain/entities"
)

type CourseRepository interface {
	Create(ctx context.Context, course *entities.Course) error
	Update(ctx context.Context, course *entities.Course) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*entities.Course, error)
	GetManyByUserID(ctx context.Context, userID int64) ([]entities.Course, error)
}
