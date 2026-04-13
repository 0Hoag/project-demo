package middleware

import pkgErrors "github.com/zeross/project-demo/pkg/errors"

var (
	errPermission    = pkgErrors.NewPermissionError(120000, "Don't have permission")
	errInvalidDevice = pkgErrors.NewHTTPError(401, "Invalid device")
)
