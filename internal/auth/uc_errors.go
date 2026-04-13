package auth

import "errors"

var (
	ErrInvalidCreds = errors.New("invalid phone or password")
)
