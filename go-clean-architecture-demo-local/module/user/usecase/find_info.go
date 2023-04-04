package usecase

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/output"
)

func (u *userUsecase) FindInfo(ctx context.Context, id string) (*output.UserInfoRes, error) {
	user, err := u.userRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	company, err := u.companyUsecase.Find(ctx, user.CompanyId)
	if err != nil {
		return nil, err
	}
	res := &output.UserInfoRes{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		CompanyInfo: *company,
	}
	return res, nil
}
