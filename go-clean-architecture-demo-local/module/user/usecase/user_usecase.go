package usecase

import (
	"context"

	_companyUsecase "github.com/jamemyjamess/go-clean-architecture-demo/module/company/usecase"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/input"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/output"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
)

// UserUsecase is a contract of business rule layer
type UserUsecase interface {
	FindInfo(ctx context.Context, id string) (*output.UserInfoRes, error)
	CreateOrUpdate(ctx context.Context, req *input.UserCreateReq) error
}

type userUsecase struct {
	// c              echo.Context
	userRepo       repository.UserRepository
	companyUsecase _companyUsecase.CompanyUsecase
}

func NewUserUsecase(usersRepo repository.UserRepository, companyUsecase _companyUsecase.CompanyUsecase) UserUsecase {
	return &userUsecase{
		// c:              c,
		userRepo:       usersRepo,
		companyUsecase: companyUsecase,
	}
}
