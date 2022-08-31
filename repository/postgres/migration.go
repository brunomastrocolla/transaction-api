package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// For migrate with files
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

type Migration struct {
	db *sqlx.DB
}

func (m Migration) Migrate(migrationDir string) error {
	driver, err := postgres.WithInstance(m.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	mg, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		return err
	}

	err = mg.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func NewMigrationRepository(db *sqlx.DB) Migration {
	return Migration{
		db: db,
	}
}
