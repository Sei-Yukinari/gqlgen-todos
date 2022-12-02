package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sei-Yukinari/gqlgen-todos/src/config"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

func Run(handler *gin.Engine) {
	port := config.Conf.App.Port
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Infof("connect to http://localhost:%s/ for GraphQL playground", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Warn("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Conf.App.TimeoutToGracefulShutdownMs)*time.Millisecond)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown:%v\n", err)
	}

	logger.Info("Server exiting")
}
