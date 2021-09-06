package services

import "errors"

type User struct {
	Name  string
	Login string
	Pass  string
}

type UserService interface {
	Login(user, pass string) (*User, error)
}

type userServiceLocal struct {
	// campos del servicio

}

type userServiceAWS struct {
	// campos del servicio

}

func (us *userServiceLocal) Login(user, pass string) (*User, error) {
	if user == "jpp" && pass == "123456" {
		return &User{"Juan Pablo", user, pass}, nil
	}
	return nil, errors.New("no se econtro el usuario")
}

func (us *userServiceAWS) Login(user, pass string) (*User, error) {
	if user == "jppAWS" && pass == "654321" {
		return &User{"Juan Pablo AWS", user, pass}, nil
	}
	return nil, errors.New("no se econtro el usuario")
}

func NewUserServiceLocal() (UserService, error) {
	return &userServiceLocal{}, nil
}

func NewUserServiceAWS() (UserService, error) {
	return &userServiceAWS{}, nil
}
