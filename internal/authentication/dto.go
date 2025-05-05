package authentication

import (
	"github.com/victorsvart/egommerce/internal/core/domain"
	"github.com/victorsvart/egommerce/pkg/rbac"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterInput) ToUser() domain.User {
	return domain.User{
		Name:     r.Name,
		Surname:  r.Surname,
		Email:    r.Email,
		Password: r.Password,
		RoleID:   rbac.UserRoleID,
	}
}
