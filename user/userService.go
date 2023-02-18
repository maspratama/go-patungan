package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUserRegister(input RegisterUserInput) (Users, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUserRegister(input RegisterUserInput) (Users, error) {
	user := Users{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return Users{}, err
	}
	user.PasswordHash = string(password)
	user.Role = "user"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
