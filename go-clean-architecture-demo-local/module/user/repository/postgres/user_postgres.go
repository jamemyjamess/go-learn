package postgres

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) repository.UserRepository {
	return &userRepository{database}
}

func (repo *userRepository) List(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}
	err := repo.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (repo *userRepository) Find(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{ID: id}
	err := repo.db.WithContext(ctx).First(&user).Error
	return user, err
}

func (repo *userRepository) Create(ctx context.Context, user *domain.User) error {
	return repo.db.WithContext(ctx).Create(user).Error
}

func (repo *userRepository) Update(ctx context.Context, user *domain.User) error {

	return repo.db.WithContext(ctx).Save(user).Error
}

func (repo *userRepository) Delete(ctx context.Context, user *domain.User) error {
	return repo.db.WithContext(ctx).Delete(user).Error
}
