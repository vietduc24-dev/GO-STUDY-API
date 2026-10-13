//go:build wireinject

package wire

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
