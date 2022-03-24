package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/liip/sheriff"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	// validate email
	_ = validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		return emailRegex.MatchString(field)

	})

	return validate

}

// ValidatorErrors func for show validation errors for each invalid fields.
// func ValidatorErrors(err error) map[string]string {
func ValidatorErrors(err error) []string {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	// Define fields map.
	// fields := map[string]string{}
	errs := []string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		// fields[err.Field()] = err.Error()
		translatedErr := fmt.Sprintf("%v", err.Translate(trans))
		errs = append(errs, translatedErr)
	}

	return errs
}

func MarshalUsers(data interface{}, groups ...string) (interface{}, error) {
	o := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(o, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
