package entities

import (
	pkg_entities "renebizelli/api/pkg/entities"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       pkg_entities.ID `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"-"`
}

func NewUser(name string, email string, password string) (*User, error) {

	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return nil, e
	}

	return &User{
		Id:       pkg_entities.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return e == nil
}
