package users

import "errors"

var wantErrors = []error{
	ErrUserNotFound,
	ErrRequiredField,
	ErrDefaultRoleUnavailable,
}

var (
	// user
	ErrUserNotFound  = errors.New("user not found")
	ErrRequiredField = errors.New("required field")

	ErrDefaultRoleUnavailable = errors.New("default role user is not configured")
)
