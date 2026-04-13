package middleware

import (
	"github.com/zeross/project-demo/internal/users"
	"github.com/zeross/project-demo/pkg/log"

	pkgCrt "github.com/zeross/project-demo/pkg/encrypter"

	"github.com/zeross/project-demo/pkg/jwt"
)

type Middleware struct {
	l           log.Logger
	userUC      users.Usecase
	jwtManager  jwt.Manager
	encrypter   pkgCrt.Encrypter
	internalKey string
}

func New(
	l log.Logger,
	userUC users.Usecase,
	jwtManager jwt.Manager,
	encrypter pkgCrt.Encrypter,
	internalKey string,
) Middleware {
	return Middleware{
		l:           l,
		userUC:      userUC,
		jwtManager:  jwtManager,
		encrypter:   encrypter,
		internalKey: internalKey,
	}
}
