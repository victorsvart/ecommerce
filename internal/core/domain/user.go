package domain

import (
	"context"
	"time"

	"github.com/victorsvart/go-ecommerce/pkg/rbac"
)

type User struct {
	ID        uint64
	Name      string `gorm:"size:255"`
	Surname   string `gorm:"size:255"`
	Email     string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	RoleID    uint
	Products  []Product
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserUseCases interface {
	GetById(ctx context.Context, id uint64) (*User, error)
	List(context.Context) ([]User, error)
	GetByEmail(context.Context, string) (*User, error)
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, uint64) error
}

type UserRepository interface {
	GetById(ctx context.Context, id uint64) (*User, error)
	List(context.Context) ([]User, error)
	GetByEmail(context.Context, string) (*User, error)
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, uint64) error
}

type UserInput struct {
	ID       *uint64 `json:"id"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func (u *UserInput) ToUser() User {
	id := uint64(0)
	if u.ID != nil {
		id = *u.ID
	}

	return User{
		ID:       id,
		Name:     u.Name,
		Surname:  u.Surname,
		Email:    u.Email,
		Password: u.Password,
		RoleID:   rbac.UserRoleID,
	}
}
