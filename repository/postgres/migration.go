package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"transaction-api/repository"
)

type Migration struct {
	db *sqlx.DB
}

func (m Migration) Migrate(migrationDir string, migrationType repository.MigrationType) error {
	driver, err := postgres.WithInstance(m.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	mg, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		return err
	}

	switch migrationType {
	case repository.MigrateUp:
		err = mg.Up()
	case repository.MigrateDown:
		err = mg.Down()
	case repository.MigrateDrop:
		err = mg.Drop()
	default:
		err = fmt.Errorf("invalid-migration-type: %s", migrationType)
	}

	if err == migrate.ErrNoChange {
		return nil
	}
	return err
}

func NewMigrationRepository(db *sqlx.DB) Migration {
	return Migration{
		db: db,
	}
}
