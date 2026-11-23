package entities

import (
	"fmt"
	"time"
)

type Course struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCourse(userID int64, title string, description string, body string) *Course {
	course := &Course{
		UserID:      userID,
		Title:       title,
		Description: description,
		Body:        body,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return course
}

func (c *Course) Validate() error {
	if c.UserID == 0 {
		return fmt.Errorf("%w: user id must not be empty", ErrValidation)
	}

	return nil
}
