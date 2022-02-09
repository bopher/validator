package validations

import (
	"github.com/bopher/utils"
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func alNumFaValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	return utils.ExtractAlphaNumPersian(v, fl.Param()) == v
}

// RegisterAlNumFaValidation register validations with translations
func RegisterAlNumFaValidation(val v.Validator) {
	val.AddValidation("alnumfa", alNumFaValidation)
	val.AddTranslation("en", "alnumfa", "Only alpha (en-fa) and numbers valid")
	val.AddTranslation("fa", "alnumfa", "فقط حروف فارسی و انگلیسی و عدد مجاز است")
}
