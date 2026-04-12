package customerrors

import (
	"errors"
	"fmt"

	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/settings"
	"github.com/go-playground/validator/v10"
)

var (
	ErrIncorrectJSON = errors.New("incorrect JSON")
	ErrParseJSON     = errors.New("failed to parse JSON")
	ErrEmptyLogin    = errors.New("user login is required")
	ErrEmptyPassword = errors.New("user password is required")

	ErrShortLogin = errors.New(fmt.Sprintf("user login is shorter than %d characters", settings.MinLoginLen))
	ErrLongLogin  = errors.New(fmt.Sprintf("user login is longer than %d characters", settings.MaxLoginLen))

	ErrShortPassword = errors.New(fmt.Sprintf("user password is shorter than %d characters", settings.MinPasswordLen))
	ErrLongPassword  = errors.New(fmt.Sprintf("user password is longer than %d characters", settings.MaxPasswordLen))
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range ve {
			field := fieldErr.Field()

			switch fieldErr.Tag() {
			case "required":
				errors[field] = "field is required"
			case "min":
				errors[field] = "too short"
			case "max":
				errors[field] = "too long"
			default:
				errors[field] = "invalid value"
			}
		}
	}

	return errors
}
