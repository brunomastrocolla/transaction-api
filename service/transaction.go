package service

import (
	"time"

	"transaction-api/entity"
	"transaction-api/repository"
)

type Transaction struct {
	transactionRepo repository.TransactionRepository
}

func (t Transaction) Create(transaction *entity.Transaction) error {
	now := time.Now()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now
	return t.transactionRepo.Create(transaction)
}

func (t Transaction) Find(id int64) (entity.Transaction, error) {
	return t.transactionRepo.Find(id)
}

func (t Transaction) FindByAccountID(id int64) ([]entity.Transaction, error) {
	return t.transactionRepo.FindByAccountID(id)
}

func (t Transaction) Update(transaction *entity.Transaction) error {
	transaction.UpdatedAt = time.Now()
	return t.transactionRepo.Update(transaction)
}

func (t Transaction) UpdateBalance(transaction *entity.Transaction) error {
	if transaction.OperationTypeID == 4 {
		transactions, err := t.FindByAccountID(transaction.AccountID)
		if err != nil {
			return err
		}

		for i := range transactions {
			current := &transactions[i]

			if transaction.Balance == 0 {
				break
			}

			switch current.OperationTypeID {
			case 1, 2, 3:
				balance := -current.Balance
				if balance == 0 {
					continue
				}
				if balance > transaction.Balance {
					current.Balance = -(balance - transaction.Balance)
					transaction.Balance = 0
				} else {
					current.Balance = 0
					transaction.Balance -= balance
				}

				if err := t.Update(current); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewTransactionService(transactionRepo repository.TransactionRepository) Transaction {
	return Transaction{
		transactionRepo: transactionRepo,
	}
}
