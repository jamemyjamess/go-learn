package handler

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/controller"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository/postgres"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/usecase"
	"github.com/jamemyjamess/go-clean-architecture-demo/pkg/database"
	"github.com/labstack/echo/v4"
)

func NewUserHandler(e *echo.Group) {
	usersRepository := postgres.NewUserRepository(database.PostgresSql)
	usersUsecase := usecase.NewUserUsecase(usersRepository, nil)
	controller.NewUsersController(e, usersUsecase)
}
