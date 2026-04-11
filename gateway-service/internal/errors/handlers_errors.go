package customerrors

import "errors"

var (
	ErrRegister = errors.New("error when registering user")
)
