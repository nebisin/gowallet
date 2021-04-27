package model

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountPayload{
		Owner:    "Test",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testRepository.CreateAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)

	return account
}

func TestRepository_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestRepository_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testRepository.GetAccount(account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestRepository_ListAccount(t *testing.T) {
	var lastAccount Account

	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testRepository.ListAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func TestRepository_UpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: 200,
	}

	account2, err := testRepository.UpdateAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestRepository_DeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testRepository.DeleteAccount(account1.ID)
	require.NoError(t, err)

	account2, err := testRepository.GetAccount(account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}