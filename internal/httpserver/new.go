package httpserver

import (
	"github.com/gin-gonic/gin"
	usersRepo "github.com/zeross/project-demo/internal/users/repository"
	pkgEncrypter "github.com/zeross/project-demo/pkg/encrypter"
	pkgLog "github.com/zeross/project-demo/pkg/log"
)

const productionMode = "production"

var ginMode = gin.DebugMode

type HTTPServer struct {
	gin          *gin.Engine
	l            pkgLog.Logger
	port         int
	db           usersRepo.DBPool
	jwtSecretKey string
	mode         string
	internalKey  string
	encrypter    pkgEncrypter.Encrypter
}

type Config struct {
	Port         int
	DB           usersRepo.DBPool
	Mode         string
	JwtSecretKey string
	InternalKey  string
	Encrypter    pkgEncrypter.Encrypter
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	if cfg.Mode == productionMode {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)

	engine := gin.Default()

	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	return &HTTPServer{
		gin:          engine,
		l:            l,
		port:         cfg.Port,
		db:           cfg.DB,
		jwtSecretKey: cfg.JwtSecretKey,
		mode:         cfg.Mode,
		internalKey:  cfg.InternalKey,
		encrypter:    cfg.Encrypter,
	}
}
