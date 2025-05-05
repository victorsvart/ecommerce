package adminhandler

import "github.com/victorsvart/egommerce/internal/core/domain"

type UserPresenter struct {
	ID       uint64 `json:"id"`
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

func ToAdminUserPresenter(u *domain.User) UserPresenter {
	return UserPresenter{
		ID:       u.ID,
		FullName: u.Name + " " + u.Surname,
		Name:     u.Name,
		Surname:  u.Surname,
		Email:    u.Email,
	}
}

func ToAdminUserPresenterSlice(u []domain.User) []UserPresenter {
	p := make([]UserPresenter, 0)
	for i := range u {
		p = append(p, ToAdminUserPresenter(&u[i]))
	}

	return p
}
