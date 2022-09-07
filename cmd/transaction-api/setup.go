package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"transaction-api/config"
	http2 "transaction-api/handler/http"
	"transaction-api/repository/postgres"
	"transaction-api/service"
)

func setupDatabase(config *config.Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	var retries int32

	for {
		retries++
		if retries > config.DatabaseMaxRetries {
			return nil, fmt.Errorf("sql-max-retries, database_url=%s, max_retries=%d",
				config.DatabaseURL, config.DatabaseMaxRetries)
		}

		db, err = sqlx.Open("postgres", config.DatabaseURL)
		if err != nil {
			zap.L().Error("sql-open-db-error", zap.Error(err))
			time.Sleep(config.DatabaseRetryDelay)
			continue
		}

		if err = db.Ping(); err != nil {
			zap.L().Error("sql-ping-db-error", zap.Error(err))
			time.Sleep(config.DatabaseRetryDelay)
			continue
		}

		break
	}

	return db, err
}

func closeDatabase(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		zap.L().Error("db-close-error", zap.Error(err))
	}
}

func setupRouter(db *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	// recover middleware
	router.Use(middleware.Recoverer)
	// zap logger middleware
	router.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			start := time.Now()
			defer func() {
				fields := []zapcore.Field{
					zap.String("method", r.Method),
					zap.String("uri", r.RequestURI),
					zap.String("remote", r.RemoteAddr),
					zap.Int("status", ww.Status()),
					zap.Int("content-length", ww.BytesWritten()),
					zap.Duration("latency", time.Since(start)),
				}
				zap.L().Info(fmt.Sprintf("%s %s", r.Method, r.RequestURI), fields...)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	})

	accountRepo := postgres.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo)
	accountHandler := http2.NewAccountHandler(accountService)
	router.Route("/accounts", func(r chi.Router) {
		r.Post("/", accountHandler.Post)
		r.Get("/{id}", accountHandler.Get)
	})

	transactionRepo := postgres.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := http2.NewTransactionHandler(transactionService)
	router.Route("/transactions", func(r chi.Router) {
		r.Post("/", transactionHandler.Post)
	})

	return router
}

func setupCli(config *config.Config) *cli.App {
	app := cli.NewApp()
	app.Name = "Pismo API"
	app.Commands = []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Starts the server",
			Action: func(c *cli.Context) error {
				return runServer(config)
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Migrate database migrate",
			Action: func(c *cli.Context) error {
				return runMigrate(config)
			},
		},
	}
	return app
}

func setupLogger(config *config.Config) error {
	var l zapcore.Level
	if err := l.Set(strings.ToUpper(config.LogLevel)); err != nil {
		return err
	}
	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(l),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "ts",
			LevelKey:   "level",
			NameKey:    "N",
			CallerKey:  "caller",
			MessageKey: "msg",
			//StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := zapConfig.Build()
	if err != nil {
		return err
	}
	_ = zap.ReplaceGlobals(logger)
	return nil
}
