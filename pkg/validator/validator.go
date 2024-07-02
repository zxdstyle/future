package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	validate = validator.New()
	en := en.New()
	uni := ut.New(en, en)
	Trans, _ = uni.GetTranslator("en")
	err := en_translations.RegisterDefaultTranslations(validate, Trans)
	validate.SetTagName("v")
	if err != nil {
		panic(err)
	}
}

func Validate(v interface{}) error {
	return validate.Struct(v)
}
