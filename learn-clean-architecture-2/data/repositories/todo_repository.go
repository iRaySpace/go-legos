package repositories

import (
	"github.com/google/uuid"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/models"
)

type (
	TodoRepository interface {
		Create(models.Todo) *models.Todo
		GetAll() []*models.Todo
	}
	LocalTodoRepository struct {
		data []*models.Todo
	}
)

func (r *LocalTodoRepository) Create(t models.Todo) *models.Todo {
	t.ID = uuid.NewString()
	r.data = append(r.data, &t)
	return &t
}

func (r *LocalTodoRepository) GetAll() []*models.Todo {
	return r.data
}
