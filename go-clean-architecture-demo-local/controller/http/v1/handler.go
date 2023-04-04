package handler

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/pkg/database"

	_companyHttp "github.com/jamemyjamess/go-clean-architecture-demo/module/company/controller"
	_companyRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/company/repository/postgres"
	_companyUsecase "github.com/jamemyjamess/go-clean-architecture-demo/module/company/usecase"
	_userHttp "github.com/jamemyjamess/go-clean-architecture-demo/module/user/controller"
	_userRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository/postgres"
	_userUsecase "github.com/jamemyjamess/go-clean-architecture-demo/module/user/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	Echo *echo.Echo
	// Cfg *configs.Configs
	Db *gorm.DB
}

func (app *App) NewRouter(e *echo.Echo) {
	e.Static("public", "assets/public")
	v1 := e.Group("/api/v1")
	//* Users group
	companyRoute := v1.Group("/company")
	companyRepository := _companyRepository.NewCompanyRepository(database.PostgresSql)
	companyUsecase := _companyUsecase.NewCompanyUsecase(companyRepository)
	_companyHttp.NewCompanyController(companyRoute, companyUsecase)

	//* Users group
	userRoute := v1.Group("/company")
	userRepository := _userRepository.NewUserRepository(database.PostgresSql)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, companyUsecase)
	_userHttp.NewUsersController(userRoute, userUsecase)

}
