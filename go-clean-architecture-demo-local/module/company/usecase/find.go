package usecase

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/entity/domain"
)

func (c *companyUsecase) Find(ctx context.Context, id string) (*domain.Company, error) {
	company, err := c.companyRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return company, nil
}
