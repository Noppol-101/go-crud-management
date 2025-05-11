package server

import (
	"context"
	"fmt"
	"log"
	"my-crud-management/internal/adapter/handler"
	"my-crud-management/internal/adapter/repository"
	"my-crud-management/internal/core/domain"
	"my-crud-management/internal/core/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sa"
	password = "p@assw0rd"
	dbname   = "mydatabase"
)

func StartServer() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbname, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logrus.Fatalln("failed to connect database")
	}

	if err := db.AutoMigrate(&domain.Brands{}, &domain.Categories{}, &domain.Products{}); err != nil {
		logrus.Fatalln("failed to auto migrate")
	}

	// REPOSITORY
	categoryRepo := repository.NewCategoryGormRepo(db)

	// SERVICE
	categoryService := service.NewCategoryService(categoryRepo)

	// HANDLER
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// ROUTES: Init router
	router, err := handler.NewRouter(handler.RouterParams{
		CategoryHandler: categoryHandler,
	})
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

	// ปิด database connection
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("failed to get sql DB from gorm: %v", err)
	} else {
		if err := sqlDB.Close(); err != nil {
			logrus.Errorf("error closing DB connection: %v", err)
		} else {
			logrus.Info("database connection closed")
		}
	}

}
