package httpserver

import (
	authHTTP "github.com/zeross/project-demo/internal/auth/delivery/http"
	authUsecase "github.com/zeross/project-demo/internal/auth/usecase"
	"github.com/zeross/project-demo/internal/middleware"
	permHTTP "github.com/zeross/project-demo/internal/permissions/delivery/http"
	permPostgres "github.com/zeross/project-demo/internal/permissions/repository/postgres"
	permUsecase "github.com/zeross/project-demo/internal/permissions/usecase"
	roleHTTP "github.com/zeross/project-demo/internal/roles/delivery/http"
	rolePostgres "github.com/zeross/project-demo/internal/roles/repository/postgres"
	roleUsecase "github.com/zeross/project-demo/internal/roles/usecase"
	userHTTP "github.com/zeross/project-demo/internal/users/delivery/http"
	userPostgres "github.com/zeross/project-demo/internal/users/repository/postgres"
	userUsecase "github.com/zeross/project-demo/internal/users/usecase"
	"github.com/zeross/project-demo/pkg/jwt"
)

func (srv HTTPServer) mapHandlers() error {
	jwtManager := jwt.NewManager(srv.jwtSecretKey)

	userRepo := userPostgres.New(srv.l, srv.db)
	roleRepo := rolePostgres.New(srv.l, srv.db)
	permRepo := permPostgres.New(srv.l, srv.db)

	userUc := userUsecase.New(srv.l, userRepo, roleRepo, srv.db)
	authUc := authUsecase.New(srv.l, jwtManager, userRepo)
	roleUc := roleUsecase.New(srv.l, roleRepo, permRepo)
	permUc := permUsecase.New(srv.l, permRepo)

	authHandler := authHTTP.New(srv.l, authUc)
	userHandler := userHTTP.New(srv.l, userUc)
	roleHandler := roleHTTP.New(srv.l, roleUc)
	permHandler := permHTTP.New(srv.l, permUc)

	// Middleware
	mw := middleware.New(srv.l, userUc, jwtManager, srv.encrypter, srv.internalKey)

	group := srv.gin.Group("/api/v1")

	authHTTP.MapRoutes(group.Group("/auth"), authHandler)
	userHTTP.MapRoutes(group.Group("/users"), userHandler, mw)
	roleHTTP.MapRoutes(group.Group("/roles"), roleHandler)
	permHTTP.MapRoutes(group.Group("/permissions"), permHandler)

	return nil
}
