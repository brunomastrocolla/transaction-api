package postgres

import (
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"
	"testing"
	"time"
	"transaction-api/entity"
)

func createTransaction(db *sqlx.DB, accountID int64, operationTypeID int32, amount, balance float64) (entity.Transaction, error) {
	transaction := entity.Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		Balance:         balance,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	transactionRepo := NewTransactionRepository(db)
	err := transactionRepo.Create(&transaction)
	return transaction, err
}

func TestTransactionRepository(t *testing.T) {

	t.Run("Test Create - Success", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("transaction_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		// create account
		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)

		// create transaction
		transaction, err := createTransaction(helper.DB, account.ID, 1, -100, -100)
		assert.NilError(t, err)
		assert.Equal(t, int64(1), transaction.ID)
	})

	t.Run("Test Find - Success", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("accounts_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		// create account
		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)

		// create transaction
		transaction, err := createTransaction(helper.DB, account.ID, 1, -100, -100)
		assert.NilError(t, err)
		assert.Equal(t, int64(1), transaction.ID)

		transactionRepo := NewTransactionRepository(helper.DB)
		transactionFromRepo, err := transactionRepo.Find(transaction.ID)
		assert.NilError(t, err)
		assert.Equal(t, transaction.ID, transactionFromRepo.ID)
		assert.Equal(t, transaction.AccountID, transactionFromRepo.AccountID)
		assert.Equal(t, transaction.OperationTypeID, transactionFromRepo.OperationTypeID)
		assert.Equal(t, transaction.Amount, transactionFromRepo.Amount)
		assert.Equal(t, transaction.Balance, transactionFromRepo.Balance)
	})

	t.Run("Test Find By Account ID - Success", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("accounts_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		// create account
		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)

		// create transaction 01
		transaction1, err := createTransaction(helper.DB, account.ID, 1, -100, -100)
		assert.NilError(t, err)
		assert.Equal(t, int64(1), transaction1.ID)

		// create transaction 02
		transaction2, err := createTransaction(helper.DB, account.ID, 1, -100, -100)
		assert.NilError(t, err)
		assert.Equal(t, int64(2), transaction2.ID)

		transactionRepo := NewTransactionRepository(helper.DB)
		transactionsFromRepo, err := transactionRepo.FindByAccountID(account.ID)
		assert.NilError(t, err)
		assert.Equal(t, 2, len(transactionsFromRepo))

		// check transaction 01
		assert.Equal(t, transaction1.ID, transactionsFromRepo[0].ID)
		assert.Equal(t, transaction1.AccountID, transactionsFromRepo[0].AccountID)
		assert.Equal(t, transaction1.OperationTypeID, transactionsFromRepo[0].OperationTypeID)
		assert.Equal(t, transaction1.Amount, transactionsFromRepo[0].Amount)
		assert.Equal(t, transaction1.Balance, transactionsFromRepo[0].Balance)

		// check transaction 02
		assert.Equal(t, transaction2.ID, transactionsFromRepo[1].ID)
		assert.Equal(t, transaction2.AccountID, transactionsFromRepo[1].AccountID)
		assert.Equal(t, transaction2.OperationTypeID, transactionsFromRepo[1].OperationTypeID)
		assert.Equal(t, transaction2.Amount, transactionsFromRepo[1].Amount)
		assert.Equal(t, transaction2.Balance, transactionsFromRepo[1].Balance)
	})

	t.Run("Test Update - Success", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("accounts_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		// create account
		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)

		// create transaction
		transaction, err := createTransaction(helper.DB, account.ID, 1, -100, -100)
		assert.NilError(t, err)
		assert.Equal(t, int64(1), transaction.ID)

		transaction.OperationTypeID = 2

		transactionRepo := NewTransactionRepository(helper.DB)
		assert.NilError(t, transactionRepo.Update(&transaction))
		transactionFromRepo, err := transactionRepo.Find(transaction.ID)
		assert.NilError(t, err)

		assert.Equal(t, transaction.ID, transactionFromRepo.ID)
		assert.Equal(t, transaction.AccountID, transactionFromRepo.AccountID)
		assert.Equal(t, transaction.OperationTypeID, transactionFromRepo.OperationTypeID)
		assert.Equal(t, transaction.Amount, transactionFromRepo.Amount)
		assert.Equal(t, transaction.Balance, transactionFromRepo.Balance)
	})

	t.Run("Test Create Closed DB - Error", func(t *testing.T) {
		helper := postgresTestHelper{}
		assert.NilError(t, helper.Start("transaction_test_db"))

		defer func() {
			assert.NilError(t, helper.Stop())
		}()

		// create account
		account, err := createNewAccount(helper.DB, "01234567890")
		assert.NilError(t, err)
		assert.Equal(t, int64(1), account.ID)

		assert.NilError(t, helper.DB.Close())

		// create transaction
		_, err = createTransaction(helper.DB, account.ID, 1, -100, -100)
		assert.Error(t, err, "sql: database is closed")
	})
}
