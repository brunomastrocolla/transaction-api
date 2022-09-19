package service

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"transaction-api/entity"
	"transaction-api/mocks"
)

func TestTransactionService(t *testing.T) {

	t.Run("Test Create - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		trans := entity.Transaction{
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          10,
			Balance:         -10,
		}

		transactionRepoMock.EXPECT().Create(gomock.Any()).Do(func(transaction *entity.Transaction) {
			transaction.ID = 1
			assert.Equal(t, trans.AccountID, transaction.AccountID)
			assert.Equal(t, trans.OperationTypeID, transaction.OperationTypeID)
			assert.Equal(t, trans.Amount, transaction.Amount)
			assert.Equal(t, trans.Balance, transaction.Balance)
			assert.Equal(t, false, transaction.CreatedAt.IsZero())
			assert.Equal(t, false, transaction.UpdatedAt.IsZero())
		}).Return(nil)

		err := transactionService.Create(&trans)
		assert.NilError(t, err)

	})

	t.Run("Test Find - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		trans := entity.Transaction{
			ID:              1,
			AccountID:       1,
			OperationTypeID: 1,
			Amount:          10,
			Balance:         -10,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		transactionRepoMock.EXPECT().Find(int64(1)).Return(trans, nil)

		transRepo, err := transactionService.Find(1)
		assert.NilError(t, err)
		assert.Equal(t, trans, transRepo)

	})

	t.Run("Test Find By Account ID - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		trans := []entity.Transaction{
			{
				ID:              1,
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          10,
				Balance:         -10,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
		}

		transactionRepoMock.EXPECT().FindByAccountID(int64(1)).Return(trans, nil)

		transRepo, err := transactionService.FindByAccountID(1)
		assert.NilError(t, err)
		assert.Equal(t, trans[0], transRepo[0])

	})

	t.Run("Test Update Balance 1 - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		transactionRepoMock.EXPECT().FindByAccountID(gomock.Any()).Return(
			[]entity.Transaction{
				{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -50,
					Balance:         -50,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              2,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -23.5,
					Balance:         -23.5,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              3,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -18.7,
					Balance:         -18.7,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
			}, nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(-13.5), transaction.Balance)
		}).Return(nil)

		trans := entity.Transaction{
			AccountID:       1,
			OperationTypeID: 4,
			Amount:          60,
			Balance:         60,
		}

		assert.NilError(t, transactionService.UpdateBalance(&trans))
		assert.Equal(t, float64(0), trans.Balance)
	})

	t.Run("Test Update Balance 2 - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		transactionRepoMock.EXPECT().FindByAccountID(gomock.Any()).Return(
			[]entity.Transaction{
				{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -50,
					Balance:         0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              2,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -23.5,
					Balance:         -13.5,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              3,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -18.7,
					Balance:         -18.7,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
				{
					ID:              4,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          60,
					Balance:         0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
			}, nil)

		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(nil)

		trans := entity.Transaction{
			AccountID:       1,
			OperationTypeID: 4,
			Amount:          100,
			Balance:         100,
		}

		assert.NilError(t, transactionService.UpdateBalance(&trans))
		assert.Equal(t, float64(67.8), trans.Balance)

	})

	t.Run("Test Update Balance 3 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		err := errors.New("find-error")
		transactionRepoMock.EXPECT().FindByAccountID(gomock.Any()).Return(
			[]entity.Transaction{}, err)

		assert.Error(t, transactionService.UpdateBalance(&entity.Transaction{
			OperationTypeID: 4}), err.Error())

	})

	t.Run("Test Update Balance 4 - Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		transactionRepoMock := mocks.NewMockTransactionRepository(ctrl)
		transactionService := NewTransactionService(transactionRepoMock)

		transactionRepoMock.EXPECT().FindByAccountID(gomock.Any()).Return(
			[]entity.Transaction{
				{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -50,
					Balance:         -50,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
				},
			}, nil)

		err := errors.New("update-error")
		transactionRepoMock.EXPECT().Update(gomock.Any()).Do(func(transaction *entity.Transaction) {
			assert.Equal(t, float64(0), transaction.Balance)
		}).Return(err)

		trans := entity.Transaction{
			AccountID:       1,
			OperationTypeID: 4,
			Amount:          100,
			Balance:         100,
		}

		assert.Error(t, transactionService.UpdateBalance(&trans), err.Error())

	})
}
