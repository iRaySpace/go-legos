package core

import (
	"context"
	"log"

	"github.com/irayspace/go-legos/learn-clean-architecture-2/app/auth"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/app/services"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/models"
	"github.com/irayspace/go-legos/learn-clean-architecture-2/data/repositories"
)

type Myapp struct{}

func (m *Myapp) Start() {
	localUserRepository := repositories.LocalUserRepository{}
	userService := services.UserService{
		UserRepository: &localUserRepository,
	}

	adminUser, _ := userService.Create(models.User{
		Username: "Administrator",
	})
	ctx := context.Background()
	ctx = context.WithValue(ctx, auth.UserIDKey{}, adminUser.ID)

	localTodoRepository := repositories.LocalTodoRepository{}
	todoService := services.TodoService{
		TodoRepository: &localTodoRepository,
		UserRepository: &localUserRepository,
	}

	todo, err := todoService.Create(ctx, models.Todo{
		Name: "Make a food",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("created", todo)

	// Guest User
	guestUser, _ := userService.Create(models.User{
		Username: "Guest",
	})
	ctx = context.WithValue(ctx, auth.UserIDKey{}, guestUser.ID)

	_, err = todoService.Create(ctx, models.Todo{
		Name: "Hello, food",
	})
	if err != nil {
		log.Println(err)
	}

	// Another Administrator
	_, err = userService.Create(models.User{
		Username: "Administrator",
	})
	if err != nil {
		log.Println(err)
	}
}
