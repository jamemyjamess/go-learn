package validatorConfig

import (
	"github.com/go-playground/validator/v10"
)

func RegisCompany(validate *validator.Validate) {
	_ = validate.RegisterValidation("requireIfAuthorizeTypeIdNotEmpty", RequireIfAuthorizeTypeIdNotEmpty)
	_ = validate.RegisterValidation("requireIfManagerIdEmpty", RequireIfManagerIdEmpty)
	_ = validate.RegisterValidation("requireIfManagerPassportEmpty", RequireIfManagerPassportEmpty)
	_ = validate.RegisterValidation("requireIfAuthorizeManagerPassportEmpty", RequireIfAuthorizeManagerPassportEmpty)
	_ = validate.RegisterValidation("requireIfAuthorizeManagerIdEmpty", RequireIfAuthorizeManagerIdEmpty)
}

func AuthorizeTypeIdNotEmpty(fl validator.FieldLevel) bool {
	for i := 0; i < fl.Top().Elem().NumField(); i++ {
		if fl.Top().Elem().Type().Field(i).Name == "AuthorizeTypeId" {
			if fl.Top().Elem().Field(i).String() != "" {
				return true
			}
		}

	}

	return false
}

func RequireIfAuthorizeTypeIdNotEmpty(fl validator.FieldLevel) bool {
	for i := 0; i < fl.Top().Elem().NumField(); i++ {
		if fl.Top().Elem().Type().Field(i).Name == "AuthorizeTypeId" {
			if fl.Field().String() == "" && AuthorizeTypeIdNotEmpty(fl) {
				return false
			}
		}

	}

	return true
}

func RequireIfManagerPassportEmpty(fl validator.FieldLevel) bool {
	for i := 0; i < fl.Top().Elem().NumField(); i++ {
		if fl.Top().Elem().Type().Field(i).Name == "ManagerPassport" {
			if fl.Field().String() == "" && fl.Top().Elem().Field(i).String() == "" {
				return false
			}
		}

	}

	return true
}

func RequireIfManagerIdEmpty(fl validator.FieldLevel) bool {
	for i := 0; i < fl.Top().Elem().NumField(); i++ {
		if fl.Top().Elem().Type().Field(i).Name == "ManagerId" {
			if fl.Field().String() == "" && fl.Top().Elem().Field(i).String() == "" {
				return false
			}
		}

	}

	return true
}

func RequireIfAuthorizeManagerPassportEmpty(fl validator.FieldLevel) bool {
	if AuthorizeTypeIdNotEmpty(fl) == false {
		return true
	}

	for i := 0; i < fl.Top().Elem().NumField(); i++ {
		if fl.Top().Elem().Type().Field(i).Name == "AuthorizeManagerPassport" {
			if fl.Field().String() == "" && fl.Top().Elem().Field(i).String() == "" {
				return false
			}
		}

	}

	return true
}

func RequireIfAuthorizeManagerIdEmpty(fl validator.FieldLevel) bool {
	if AuthorizeTypeIdNotEmpty(fl) == false {
		return true
	}

	for i := 0; i < fl.Top().Elem().NumField(); i++ {
		if fl.Top().Elem().Type().Field(i).Name == "AuthorizeManagerId" {
			if fl.Field().String() == "" && fl.Top().Elem().Field(i).String() == "" {
				return false
			}
		}

	}

	return true
}
