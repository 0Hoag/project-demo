package http

import (
	"github.com/zeross/project-demo/internal/permissions"
	pkgErrors "github.com/zeross/project-demo/pkg/errors"
)

var (
	errWrongQuery = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody  = pkgErrors.NewHTTPError(140003, "Wrong body")

	errPermissionNotFound = pkgErrors.NewHTTPError(145005, "Permission not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case permissions.ErrPermissionNotFound:
		return errPermissionNotFound
	default:
		panic(err)
	}
}
