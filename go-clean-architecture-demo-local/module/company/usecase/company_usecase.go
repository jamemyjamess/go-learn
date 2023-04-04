package usecase

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/entity/domain"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/repository"
)

// CompanyUsecase is a contract of business rule layer
type CompanyUsecase interface {
	Find(ctx context.Context, id string) (*domain.Company, error)
}

type companyUsecase struct {
	companyRepo repository.CompanyRepository
}

func NewCompanyUsecase(companyRepo repository.CompanyRepository) CompanyUsecase {
	return &companyUsecase{
		companyRepo: companyRepo,
	}
}
