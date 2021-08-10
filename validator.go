package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator interface
type Validator interface {
	// Validator get original validator instance
	Validator() *validator.Validate
	// AddValidation add new validator
	AddValidation(tag string, v validator.Func)

	// AddTranslation register new translation message to validator translator
	AddTranslation(locale string, key string, message string)
	// Translate generate translation
	Translate(locale string, key string, placeholders map[string]string) string
	// TranslateStruct generate translation for struct
	TranslateStruct(s interface{}, locale string, key string, field string, placeholders map[string]string) string

	// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
	// Return translated errors list on fails
	// Return nil on no error
	// use vTitle tag for define field title, (use vTitle_locale) for localization
	// use vParam tag for define param title in multifield validations (like eqcsfield). this tag follow vTitle tag rules
	// use vFormat tag for format parameter as number
	StructLocale(locale string, s interface{}) ErrorResponse
	// Validate using default locale
	// @see StructLocale
	Struct(s interface{}) ErrorResponse
	// StructExcept validates all fields except the ones passed in.
	// Return translated errors list on fails
	// Return nil on no error
	StructExceptLocale(locale string, s interface{}, fields ...string) ErrorResponse
	// Validate using default locale
	// @see StructExceptLocale
	StructExcept(s interface{}, fields ...string) ErrorResponse
	// StructPartial validates the fields passed in only, ignoring all others.
	// Return translated errors list on fails
	// Return nil on no error
	StructPartialLocale(locale string, s interface{}, fields ...string) ErrorResponse
	// Validate using default locale
	// @see StructPartialLocale
	StructPartial(s interface{}, fields ...string) ErrorResponse

	// Var validates a single variable using tag style validation.
	// Return translated errors list on fails
	// Return nil on no error
	VarLocale(locale string, params ValidatorParam, field interface{}, tag string, messages map[string]string) ErrorResponse
	// Validate using default locale
	// @see VarLocale
	Var(params ValidatorParam, field interface{}, tag string, messages map[string]string) ErrorResponse
	// VarWithValue validates a single variable, against another variable/field's value using tag style validation
	// Return translated errors list on fails
	// Return nil on no error
	VarWithValueLocale(locale string, params ValidatorParam, field interface{}, other interface{}, tag string, messages map[string]string) ErrorResponse
	// Validate using default locale
	// @see VarWithValueLocale
	VarWithValue(params ValidatorParam, field interface{}, other interface{}, tag string, messages map[string]string) ErrorResponse
}
