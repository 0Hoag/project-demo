package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (srv HTTPServer) Run() error {
	err := srv.mapHandlers()
	if err != nil {
		return err
	}

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", srv.port),
		Handler: srv.gin,
	}

	go func() {
		srv.l.Infof(context.Background(), "Started server on :%d", srv.port)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.l.Fatalf(context.Background(), "listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	srv.l.Infof(context.Background(), "Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		srv.l.Errorf(context.Background(), "Server forced to shutdown: %v", err)
		return err
	}

	srv.l.Infof(context.Background(), "Server exiting")
	return nil
}
