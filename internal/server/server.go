package server

import (
	"context"
	"fmt"
	"my-crud-management/internal/adapter/handler"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartServer() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// ROUTES: Init router
	router, err := handler.NewRouter(handler.RouterParams{})
	if err != nil {
		logrus.Fatalf("failed to initialize router : %v", err)
	}

	// SERVER: Start server
	listenAddress := fmt.Sprintf(":%v", viper.GetInt("app.server.port"))
	logrus.Infof("starting server on %s", listenAddress)

	go func() {
		if err := router.Serve(listenAddress); err != nil {
			logrus.Fatalf("error starting HTTP server: %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("shutdown signal received")

	if err := router.ShutdownWithTimeout(10 * time.Second); err != nil {
		logrus.Errorf("error shutting down server: %v", err)
	} else {
		logrus.Info("server shutdown completed")
	}

}
