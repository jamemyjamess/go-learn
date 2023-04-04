package controller

import (
	"log"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/input"
	"github.com/labstack/echo/v4"
)

func (userController *userController) FindInfo(c echo.Context) error {
	req := &input.UserCreateReq{}
	if err := c.Bind(req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	res, err := userController.UserUsecase.FindInfo(ctx, req.ID)
	if err != nil {
		return err
	}
	log.Println("res:", res)

	return nil
}
