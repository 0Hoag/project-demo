package http

import (
	"github.com/zeross/project-demo/internal/auth"
	pkgErrors "github.com/zeross/project-demo/pkg/errors"
)

var (
	errWrongBody    = pkgErrors.NewHTTPError(140003, "Wrong body")
	errInvalidCreds = pkgErrors.NewHTTPError(140401, "Invalid phone or password")
)

func (h handler) mapError(err error) error {
	switch err {
	case auth.ErrInvalidCreds:
		return errInvalidCreds
	default:
		panic(err)
	}
}
