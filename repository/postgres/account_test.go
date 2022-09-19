package postgres

import (
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"
	"testing"
	"time"
	"transaction-api/entity"
)

func createNewAccount(db *sqlx.DB, documentNumber string) (entity.Account, error) {
	// create account
	account := entity.Account{
		DocumentNumber: documentNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	accountRepo := NewAccountRepository(db)
	err := accountRepo.Create(&account)
	return account, err
}

func TestAccountRepository(t *testing.T) {

	t.Run("Test Create - Success", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("account_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)
	})

	t.Run("Test Find - Success", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("accounts_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)

		accountRepo := NewAccountRepository(helper.DB)
		accountFromRepo, err := accountRepo.Find(account.ID)
		assert.NilError(t, err)
		assert.Equal(t, account.ID, accountFromRepo.ID)
		assert.Equal(t, account.DocumentNumber, accountFromRepo.DocumentNumber)
	})

	t.Run("Test Create Closed DB - Error", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("account_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		assert.NilError(t, helper.DB.Close())
		_, err := createNewAccount(helper.DB, "01234567890")
		assert.Error(t, err, "sql: database is closed")
	})

}
