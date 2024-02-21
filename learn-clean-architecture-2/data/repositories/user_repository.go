package repositories

import (
	"github.com/google/uuid"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/models"
)

type (
	UserRepository interface {
		Create(models.User) *models.User
		FindByID(id string) *models.User
		FindByUsername(username string) *models.User
	}
	LocalUserRepository struct {
		data []*models.User
	}
)

func (r *LocalUserRepository) Create(u models.User) *models.User {
	u.ID = uuid.NewString()
	r.data = append(r.data, &u)
	return &u
}

func (r *LocalUserRepository) FindByID(id string) *models.User {
	for _, s := range r.data {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func (r *LocalUserRepository) FindByUsername(username string) *models.User {
	for _, s := range r.data {
		if s.Username == username {
			return s
		}
	}
	return nil
}
