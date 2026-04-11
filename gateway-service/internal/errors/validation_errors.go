package customerrors

import (
	"errors"
	"fmt"

	"github.com/NikitaTumanov/ai-editor-platform/gateway-service/internal/settings"
)

var (
	ErrIncorrectJSON = errors.New("incorrect JSON")
	ErrEmptyLogin    = errors.New("user login is required")
	ErrEmptyPassword = errors.New("user password is required")

	ErrShortLogin = errors.New(fmt.Sprintf("user login is shorter than %d characters", settings.MinLoginLen))
	ErrLongLogin  = errors.New(fmt.Sprintf("user login is longer than %d characters", settings.MaxLoginLen))

	ErrShortPassword = errors.New(fmt.Sprintf("user password is shorter than %d characters", settings.MinPasswordLen))
	ErrLongPassword  = errors.New(fmt.Sprintf("user password is longer than %d characters", settings.MaxPasswordLen))
)
