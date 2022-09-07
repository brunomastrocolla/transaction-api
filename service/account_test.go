package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"transaction-api/entity"
	"transaction-api/mocks"
)

func TestAccountService(t *testing.T) {

	t.Run("Test Create - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepoMock := mocks.NewMockAccountRepository(ctrl)
		accountService := NewAccountService(accountRepoMock)

		acc := entity.Account{
			ID:             1,
			DocumentNumber: "01234567890",
		}

		accountRepoMock.EXPECT().Create(gomock.Any()).Do(func(account *entity.Account) {
			assert.Equal(t, acc.ID, account.ID)
			assert.Equal(t, acc.DocumentNumber, account.DocumentNumber)
			assert.Equal(t, false, account.CreatedAt.IsZero())
			assert.Equal(t, false, account.UpdatedAt.IsZero())
		}).Return(nil)

		assert.NilError(t, accountService.Create(&acc))

	})

	t.Run("Test Find - Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepoMock := mocks.NewMockAccountRepository(ctrl)
		accountService := NewAccountService(accountRepoMock)

		acc := entity.Account{
			ID:             1,
			DocumentNumber: "01234567890",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		accountRepoMock.EXPECT().Find(int64(1)).Return(acc, nil)

		repoAccount, err := accountService.Find(1)
		assert.NilError(t, err)
		assert.Equal(t, acc.ID, repoAccount.ID)
		assert.Equal(t, acc.DocumentNumber, repoAccount.DocumentNumber)
		assert.Equal(t, acc.CreatedAt, repoAccount.CreatedAt)
		assert.Equal(t, acc.UpdatedAt, repoAccount.UpdatedAt)

	})

}
