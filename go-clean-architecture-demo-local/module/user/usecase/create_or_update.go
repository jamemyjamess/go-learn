package usecase

import (
	"context"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/input"
)

func (u *userUsecase) CreateOrUpdate(ctx context.Context, req *input.UserCreateReq) error {
	// if err := u.c.Validate(req); err != nil {
	// 	return err
	// }
	user := input.CreateReqToUserDomain(req)
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
