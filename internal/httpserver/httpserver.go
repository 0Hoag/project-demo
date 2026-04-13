package httpserver

import (
	"context"
	"fmt"
)

func (srv HTTPServer) Run() error {
	err := srv.mapHandlers()
	if err != nil {
		return err
	}

	srv.l.Infof(context.Background(), "Started server on :%d", srv.port)
	return srv.gin.Run(fmt.Sprintf(":%d", srv.port))
}
