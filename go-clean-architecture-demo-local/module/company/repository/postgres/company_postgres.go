package postgres

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/entity/domain"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/company/repository"
	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(database *gorm.DB) repository.CompanyRepository {
	return &companyRepository{database}
}

func (repo *companyRepository) Find(ctx context.Context, id string) (*domain.Company, error) {
	user := &domain.Company{ID: id}
	err := repo.db.WithContext(ctx).First(&user).Error
	return user, err
}
