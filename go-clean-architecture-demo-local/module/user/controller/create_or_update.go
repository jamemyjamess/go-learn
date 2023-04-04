package controller

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/input"
	"github.com/labstack/echo/v4"
)

func (userController *userController) CreateOrUpdate(c echo.Context) error {
	req := &input.UserCreateReq{}
	if err := c.Bind(req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	err := userController.UserUsecase.CreateOrUpdate(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
