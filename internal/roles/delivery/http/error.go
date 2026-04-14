package http

import (
	"github.com/zeross/project-demo/internal/permissions"
	"github.com/zeross/project-demo/internal/roles"
	pkgErrors "github.com/zeross/project-demo/pkg/errors"
)

var (
	errWrongQuery = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody  = pkgErrors.NewHTTPError(140003, "Wrong body")

	errRoleNotFound            = pkgErrors.NewHTTPError(144005, "Role not found")
	errPermissionNotFound      = pkgErrors.NewHTTPError(145005, "Permission not found")
	errPermissionAlreadyLinked = pkgErrors.NewHTTPError(144009, "Permission already linked to this role")
)

func (h handler) mapError(err error) error {
	switch err {
	case roles.ErrRoleNotFound:
		return errRoleNotFound
	case permissions.ErrPermissionNotFound:
		return errPermissionNotFound
	case roles.ErrPermissionAlreadyLinked:
		return errPermissionAlreadyLinked
	default:
		return err
	}
}
