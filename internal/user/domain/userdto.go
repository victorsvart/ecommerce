package domain

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
	}
}
