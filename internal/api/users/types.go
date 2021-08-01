package users

import "github.com/kashyaprahul94/go-boilerplate/pkg/users"

type CreateUserRequestPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UpdateUserRequestPayload struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (u *CreateUserRequestPayload) ToUser() *users.User {
	return users.NewUser(0, u.FirstName, u.LastName, u.Email)
}

func (u *UpdateUserRequestPayload) ToUser() *users.User {
	return users.NewUser(u.ID, u.FirstName, u.LastName, u.Email)
}
