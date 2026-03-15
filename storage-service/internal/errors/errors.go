package customerrors

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrDocumentNotFound = errors.New("document not found")
)
