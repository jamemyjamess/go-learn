package repository

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/entity/domain"
)

type CompanyRepository interface {
	Find(ctx context.Context, id string) (*domain.Company, error)
}
