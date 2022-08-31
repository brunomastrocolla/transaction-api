package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transaction-api/internal/config"
	"transaction-api/internal/repository/postgres"

	"go.uber.org/zap"
)

func runServer(config *config.Config) error {
	db, err := setupDatabase(config)
	if err != nil {
		zap.L().Fatal("setup-database-error", zap.Error(err))
	}
	defer closeDatabase(db)

	router := setupRouter(db)
	server := &http.Server{
		Addr:    config.HTTPServerAddress,
		Handler: router,
	}
	shutdown := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			zap.L().Error("server-shutdown-error", zap.Error(err))
		}
		zap.L().Info("server-shutdown-success")
		close(shutdown)
	}()

	zap.L().Info("server-listen-and-serve")
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	<-shutdown
	return nil
}

func runMigrate(config *config.Config) error {
	db, err := setupDatabase(config)
	if err != nil {
		zap.L().Fatal("setup-database-error", zap.Error(err))
	}
	defer closeDatabase(db)

	zap.L().Info("db-migrate-started", zap.String("migration-dir", config.DatabaseMigrationDir))

	migrate := postgres.NewMigrationRepository(db)
	if err := migrate.Migrate(config.DatabaseMigrationDir); err != nil {
		zap.L().Error("db-migrate-error", zap.Error(err))
		return err
	}

	zap.L().Info("db-migrate-finished")
	return nil
}

func main() {
	config := config.NewConfig()

	if err := setupLogger(&config); err != nil {
		zap.L().Fatal("setup-logger-error", zap.Error(err))
	}

	cli := setupCli(&config)
	if err := cli.Run(os.Args); err != nil {
		zap.L().Fatal("run-cli-error", zap.Error(err))
	}
}
