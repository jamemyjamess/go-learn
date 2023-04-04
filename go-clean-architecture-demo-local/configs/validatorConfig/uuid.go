package validatorConfig

import (
	"strings"

	"github.com/jamemyjamess/go-clean-architecture-demo/pkg/utils"

	"github.com/go-playground/validator/v10"
)

func RegisUUIDValidate(validate *validator.Validate) {
	validate.RegisterValidation("uuid_format_if_not_empty", ValidateUUIDFormatIfNotEmpty)
}

func ValidateUUIDFormatIfNotEmpty(fl validator.FieldLevel) bool {
	if strings.TrimSpace(fl.Field().String()) == "" {
		return true
	}
	return utils.IsValidUUID(fl.Field().String())
}
