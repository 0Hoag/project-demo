package http

import (
	"github.com/zeross/project-demo/internal/users"
	pkgErrors "github.com/zeross/project-demo/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(140001, "Wrong pagination query")
	errWrongQuery           = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody            = pkgErrors.NewHTTPError(140003, "Wrong body")

	// User errors
	errUserNotFound           = pkgErrors.NewHTTPError(143005, "User not found")
	errDefaultRoleUnavailable = pkgErrors.NewHTTPError(143006, "Default role is not configured; run DB migrations")
)

func (h handler) mapError(err error) error {
	switch err {
	case users.ErrUserNotFound:
		return errUserNotFound
	case users.ErrDefaultRoleUnavailable:
		return errDefaultRoleUnavailable
	default:
		return err
	}
}
