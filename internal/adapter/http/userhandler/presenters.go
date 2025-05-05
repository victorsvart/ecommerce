package userhandler

import "github.com/victorsvart/egommerce/internal/core/domain"

type UserPresenter struct {
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
}

func ToUserPresenter(u *domain.User) UserPresenter {
	return UserPresenter{
		FullName: u.Name + " " + u.Surname,
		Name:     u.Name,
		Surname:  u.Surname,
		Email:    u.Email,
		Contact:  u.Contact,
	}
}

func ToUserPresenterSlice(u []domain.User) []UserPresenter {
	p := make([]UserPresenter, 0)
	for i := range u {
		p = append(p, ToUserPresenter(&u[i]))
	}

	return p
}
