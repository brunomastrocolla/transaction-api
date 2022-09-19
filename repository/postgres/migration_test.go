package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"
	"testing"
	"transaction-api/repository"
)

func TestMigrationRepository(t *testing.T) {

	t.Run("Test Migrate Up - Success", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
			if db != nil {
				assert.NilError(t, db.Close())
			}
		}()

		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate(postgresMigrationDir, repository.MigrateUp)
		assert.NilError(t, err)
	})

	t.Run("Test Migrate Up (No Changes) - Success", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
			if db != nil {
				assert.NilError(t, db.Close())
			}
		}()

		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate(postgresMigrationDir, repository.MigrateUp)
		assert.NilError(t, err)
		err = migrateRepo.Migrate(postgresMigrationDir, repository.MigrateUp)
		assert.NilError(t, err)
	})

	t.Run("Test Migrate Down - Success", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
			if db != nil {
				assert.NilError(t, db.Close())
			}
		}()

		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate(postgresMigrationDir, repository.MigrateUp)
		assert.NilError(t, err)
		err = migrateRepo.Migrate(postgresMigrationDir, repository.MigrateDown)
		assert.NilError(t, err)
	})

	t.Run("Test Migrate Drop - Success", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
			if db != nil {
				assert.NilError(t, db.Close())
			}
		}()

		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate(postgresMigrationDir, repository.MigrateDrop)
		assert.NilError(t, err)
	})

	t.Run("Test Migrate Invalid Type - Error", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
			if db != nil {
				assert.NilError(t, db.Close())
			}
		}()

		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate(postgresMigrationDir, "invalid")
		assert.Error(t, err, "invalid-migration-type: invalid")
	})

	t.Run("Test Migrate Closed DB - Error", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
		}()

		assert.NilError(t, db.Close())
		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate(postgresMigrationDir, "")
		assert.Error(t, err, "sql: database is closed")
	})

	t.Run("Test Migrate Invalid Migration Dir - Error", func(t *testing.T) {
		driverURL, terminate, err := newPostgresContainer("migrate_test_db")
		assert.NilError(t, err)

		db, err := sqlx.Open("postgres", driverURL)
		assert.NilError(t, err)

		defer func() {
			if terminate != nil {
				assert.NilError(t, terminate(context.Background()))
			}
			if db != nil {
				assert.NilError(t, db.Close())
			}
		}()

		migrateRepo := NewMigrationRepository(db)
		err = migrateRepo.Migrate("", "")
		assert.Error(t, err, "URL cannot be empty")
	})
}
