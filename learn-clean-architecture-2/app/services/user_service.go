package services

import (
	"errors"

	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/models"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (us *UserService) Create(u models.User) (*models.User, error) {
	user := us.UserRepository.FindByUsername(u.Username)
	if user != nil {
		return nil, errors.New("unable to create User, already created")
	}
	return us.UserRepository.Create(u), nil
}

func (us *UserService) FindByID(id string) *models.User {
	return us.UserRepository.FindByID(id)
}
