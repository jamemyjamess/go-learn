package output

import (
	"time"

	_companyDomain "github.com/jamemyjamess/go-clean-architecture-demo/module/company/entity/domain"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
)

type UserCreateRes struct {
	ID        string    `json:"id" xml:"id"`
	Name      string    `json:"name" xml:"name"`
	Email     string    `json:"email" xml:"email"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at"`
}

type UserInfoRes struct {
	ID          string                 `json:"id" xml:"id"`
	Name        string                 `json:"name" xml:"name"`
	Email       string                 `json:"email" xml:"email"`
	CreatedAt   time.Time              `json:"created_at" xml:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" xml:"updated_at"`
	CompanyInfo _companyDomain.Company `json:"company_info" xml:"company_info"`
}

func UserToCreateRes(company *domain.User) *UserCreateRes {
	return &UserCreateRes{
		ID:   company.ID,
		Name: company.Name,
	}
}
