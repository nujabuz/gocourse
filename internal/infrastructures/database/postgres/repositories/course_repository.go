package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nujabuz/gocourse/internal/domain/entities"
	"github.com/nujabuz/gocourse/internal/domain/repositories"
)

type courseRepository struct {
	db *pgxpool.Pool
}

func NewCourseRepository(db *pgxpool.Pool) repositories.CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (repo *courseRepository) Create(ctx context.Context, course *entities.Course) error {
	query := "INSERT INTO courses (user_id, title, description, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := repo.db.Exec(ctx, query, course.UserID, course.Title, course.Description, course.Body, course.CreatedAt, course.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *courseRepository) Update(ctx context.Context, course *entities.Course) error {
	query := "UPDATE courses SET title = $2, description = $3, body = $4, updated_at = $5 WHERE id = $1;"

	_, err := repo.db.Exec(ctx, query, course.ID, course.Title, course.Description, course.Body, course.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *courseRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM courses WHERE id = $1;"

	_, err := repo.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *courseRepository) GetByID(ctx context.Context, id int64) (*entities.Course, error) {
	query := "SELECT id, user_id, title, description, body, created_at, updated_at FROM courses WHERE id = $1;"

	var course entities.Course
	if err := repo.db.QueryRow(ctx, query, id).Scan(&course.ID, &course.UserID, &course.Title, &course.Description, &course.Body, &course.CreatedAt, &course.UpdatedAt); err != nil {
		return nil, err
	}

	return &course, nil
}

func (repo *courseRepository) GetManyByUserID(ctx context.Context, userID int64) ([]entities.Course, error) {
	query := "SELECT id, user_id, title, description, body, created_at, updated_at FROM courses WHERE user_id = $1;"

	rows, err := repo.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []entities.Course
	for rows.Next() {
		var id int64
		var userID int64
		var title string
		var description string
		var body string
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(&id, &userID, &title, &description, &body, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		courses = append(courses, entities.Course{
			ID:          id,
			UserID:      userID,
			Title:       title,
			Description: description,
			Body:        body,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}

	return courses, nil
}
