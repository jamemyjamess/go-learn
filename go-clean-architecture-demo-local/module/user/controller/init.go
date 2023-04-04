package controller

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/usecase"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	FindInfo(c echo.Context) error
	CreateOrUpdate(c echo.Context) error
}

type userController struct {
	UserUsecase usecase.UserUsecase
}

func NewUsersController(e *echo.Group, userUsecase usecase.UserUsecase) {
	userControllers := &userController{
		UserUsecase: userUsecase,
	}
	e.Group("/user")
	{
		e.GET("/:id", userControllers.FindInfo)
		e.PUT("/:id", userControllers.CreateOrUpdate)
	}
}
