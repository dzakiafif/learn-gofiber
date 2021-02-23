package config

import (
    "log"
    "github.com/go-playground/locales/id"
	"github.com/go-playground/validator/v10"
    ut "github.com/go-playground/universal-translator"
    id_translate "github.com/go-playground/validator/v10/translations/id"
)

type ErrorResponse struct {
    Value       string
}


func ValidateStruct(entity interface{}) []string {

    var errors []string

    translate := id.New()
    universal := ut.New(translate,translate)
    trans,notfound := universal.GetTranslator("id")

    if !notfound {
        log.Fatal("translator ora di temukno. ceken maneh")
    }

    validate := validator.New()

    if errorValidate := id_translate.RegisterDefaultTranslations(validate,trans); errorValidate != nil {
        log.Fatal(errorValidate)
    }

    _ = validate.RegisterTranslation("required", trans, func(universalTranslate ut.Translator) error {
		return universalTranslate.Add("required", "{0} dikandani required kok pancet ngeyel kosong ae", true)
	}, func(universalTranslate ut.Translator, errorField validator.FieldError) string {
		result, _ := universalTranslate.T("required", errorField.Field())
		return result
	})

    err := validate.Struct(entity)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            errors = append(errors,err.Translate(trans))
        }
    }
    return errors
}