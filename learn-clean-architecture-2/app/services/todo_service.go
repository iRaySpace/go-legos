package services

import (
	"context"
	"errors"

	"github.com/irayspace/go-legos/learn-clean-architecture-2/app/auth"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/models"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/repositories"
)

type TodoService struct {
	TodoRepository repositories.TodoRepository
	UserRepository repositories.UserRepository
}

func (ts *TodoService) Create(ctx context.Context, t models.Todo) (*models.Todo, error) {
	userId := ctx.Value(auth.UserIDKey{}).(string)
	existingUser := ts.UserRepository.FindByID(userId)
	if existingUser.Username != "Administrator" {
		return nil, errors.New("unable to create Todo, user is not Administrator")
	}
	return ts.TodoRepository.Create(t), nil
}
