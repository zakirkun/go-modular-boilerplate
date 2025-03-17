package validator

import "github.com/go-playground/validator"

// CustomValidator is a custom validator for Echo
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

// Validate validates a struct
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
