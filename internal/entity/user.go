package entity

import (
	"errors"
	"net/http"

	customerror "github.com/krishnapramodaradhi/xpressbuy-api/internal/util/customError"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

func NewUser(id string, firstName string, lastName string, email string) *User {
	return &User{Id: id, FirstName: firstName, LastName: lastName, Email: email}
}

func (u *User) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return customerror.New(http.StatusBadRequest, "Auth Failed")
		}
		return customerror.New(http.StatusInternalServerError, err.Error())
	}
	return nil
}
