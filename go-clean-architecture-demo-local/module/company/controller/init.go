package controller

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/usecase"
	"github.com/labstack/echo/v4"
)

type CompanyController interface {
	Find(c echo.Context) error
	CreateOrUpdate(c echo.Context) error
}

type companyController struct {
	CompanyUsecase usecase.CompanyUsecase
}

func NewCompanyController(e *echo.Group, companyUsecase usecase.CompanyUsecase) {
	userControllers := &companyController{
		CompanyUsecase: companyUsecase,
	}
	e.Group("/company")
	{
		e.GET("/:id", userControllers.Find)
	}
}
