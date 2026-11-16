// Package entities
package entities

import (
	"fmt"
	"time"
)

const (
	MaxUsername = 16
	MinUsername = 3
	MaxPassword = 32
	MinPassword = 8
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username string, password string) *User {
	return &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) Validate() error {
	if err := u.validateUsername(); err != nil {
		return err
	}

	if err := u.validatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *User) validateUsername() error {
	if u.Username == "" {
		return fmt.Errorf("%w: username must not be empty", ErrValidation)
	}

	if len(u.Username) > MaxUsername {
		return fmt.Errorf("%w: username length must not be more than %d characters", ErrValidation, MaxUsername)
	}

	if len(u.Username) < MinUsername {
		return fmt.Errorf("%w: username length must not be less than %d characters", ErrValidation, MinUsername)
	}

	return nil
}

func (u *User) validatePassword() error {
	if u.Password == "" {
		return fmt.Errorf("%w: password must not be empty", ErrValidation)
	}

	if len(u.Password) > MaxPassword {
		return fmt.Errorf("%w: password length must not be more than %d characters", ErrValidation, MaxUsername)
	}

	if len(u.Password) < MinPassword {
		return fmt.Errorf("%w: password length must not be less than %d characters", ErrValidation, MinUsername)
	}

	return nil
}
