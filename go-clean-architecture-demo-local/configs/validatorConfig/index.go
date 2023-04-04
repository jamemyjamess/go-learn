package validatorConfig

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}
	// translate all error at once
	validatorErr := err.(validator.ValidationErrors)
	return errors.New(fmt.Sprintf("%s", validatorErr.Translate(cv.trans)))
}

func Init(e *echo.Echo) {
	validate := validator.New()
	enTranslator := en.New()
	uni := ut.New(enTranslator, enTranslator)
	trans, _ := uni.GetTranslator("en")
	err := enTranslations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	RegisCompany(validate)
	RegisUUIDValidate(validate)

	e.Validator = &CustomValidator{validator: validate, trans: trans}
}
