package service

import "errors"

var ErrUserAlreadyExists = errors.New("user exists")

type UserService interface {
	CreateUser(name, login, password string) error
}
