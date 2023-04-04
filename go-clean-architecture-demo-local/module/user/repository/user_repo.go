package repository

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
)

// UserRepository is a contract of database connection adapter layer
type UserRepository interface {
	List(ctx context.Context) ([]domain.User, error)
	Find(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, user *domain.User) error
}
