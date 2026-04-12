package customerrors

import "errors"

var (
	ErrRegister           = errors.New("failed to register user")
	ErrLogin              = errors.New("failed to login")
	ErrQuestion           = errors.New("failed to validate question")
	ErrUpdateDocumentByID = errors.New("failed to update document by id")
)
