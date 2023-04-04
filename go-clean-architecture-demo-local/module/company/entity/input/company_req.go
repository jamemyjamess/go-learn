package input

import (
	"time"

	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
)

type CompanyCreateReq struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @Name UserCreateInput

func MakeTestCreateInput() (req *CompanyCreateReq) {
	return &CompanyCreateReq{
		ID:   "test",
		Name: "test",
	}
}

func CreateReqToCompanyDomain(req *CompanyCreateReq) (company *domain.User) {
	return &domain.User{
		ID:   req.ID,
		Name: req.Name,
	}
}
