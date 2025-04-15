package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint64
	Name      string `gorm:"size:255"`
	Surname   string `gorm:"size:255"`
	Email     string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserUseCases interface {
	List(context.Context) ([]User, error)
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, uint64) error
}

type UserRepository interface {
	List(context.Context) ([]User, error)
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, uint64) error
}
