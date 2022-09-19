package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"transaction-api/repository"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/thanhpk/randstr"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	postgresMigrationDir = "file://../../db/migrations"
)

type postgresTestHelper struct {
	DB        *sqlx.DB
	Terminate func(context.Context) error
}

func (h *postgresTestHelper) Start(dbName string) error {
	var driverURL string
	var err error

	if driverURL, h.Terminate, err = newPostgresContainer(dbName); err != nil {
		return err
	}
	if h.DB, err = sqlx.Open("postgres", driverURL); err != nil {
		return err
	}

	migration := NewMigrationRepository(h.DB)
	if err = migration.Migrate(postgresMigrationDir, repository.MigrateDrop); err != nil {
		return err
	}
	if err = migration.Migrate(postgresMigrationDir, repository.MigrateUp); err != nil {
		return err
	}

	return nil
}

func (h *postgresTestHelper) Stop() error {
	if h.DB != nil {
		if err := h.DB.Close(); err != nil {
			return err
		}
	}
	if h.Terminate != nil {
		if err := h.Terminate(context.Background()); err != nil {
			return err
		}
	}
	return nil
}

func newPostgresContainer(dbName string) (string, func(context.Context) error, error) {
	ctx := context.Background()
	templateURL := "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	username := randstr.String(16)
	password := randstr.String(16)

	// Up container
	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		Started: true,
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "postgres:13-alpine",
			ExposedPorts: []string{
				"0:5432",
			},
			Env: map[string]string{
				"POSTGRES_USER":     username,
				"POSTGRES_PASSWORD": password,
				"POSTGRES_DB":       dbName,
				"POSTGRES_SSL_MODE": "disable",
			},
			Cmd: []string{
				"postgres", "-c", "fsync=off",
			},
			WaitingFor: wait.ForSQL(
				"5432/tcp",
				"postgres",
				func(p nat.Port) string {
					return fmt.Sprintf(templateURL, username, password, p.Port(), dbName)
				},
			).Timeout(time.Second * 5),
		},
	})
	if err != nil {
		return "", func(context.Context) error { return nil }, err
	}

	// Find port of container
	ports, err := c.Ports(ctx)
	if err != nil {
		return "", func(context.Context) error { return nil }, err
	}

	// Format driverURL
	driverURL := fmt.Sprintf(templateURL, username, password, ports["5432/tcp"][0].HostPort, dbName)
	return driverURL, c.Terminate, nil
}
