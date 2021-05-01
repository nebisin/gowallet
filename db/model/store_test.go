package model

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSQLRepository_TransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	result, err := store.TransferTx(TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        5,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	transfer := result.Transfer
	require.NotEmpty(t, transfer)
	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.Equal(t, int64(5), transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	_, err = testRepository.GetTransfer(transfer.ID)
	require.NoError(t, err)

	fromEntry := result.FromEntry
	require.NotEmpty(t, fromEntry)
	require.Equal(t, account1.ID, fromEntry.AccountID)
	require.Equal(t, int64(-5), fromEntry.Amount)
	require.NotZero(t, fromEntry.ID)
	require.NotZero(t, fromEntry.CreatedAt)

	_, err = testRepository.GetEntry(fromEntry.ID)
	require.NoError(t, err)

	toEntry := result.ToEntry
	require.NotEmpty(t, toEntry)
	require.Equal(t, account2.ID, toEntry.AccountID)
	require.Equal(t, int64(5), toEntry.Amount)
	require.NotZero(t, toEntry.ID)
	require.NotZero(t, toEntry.CreatedAt)

	_, err = testRepository.GetEntry(toEntry.ID)
	require.NoError(t, err)

	fromAccount := result.FromAccount
	require.NotEmpty(t, fromAccount)
	require.Equal(t, account1.ID, fromAccount.ID)

	toAccount := result.ToAccount
	require.NotEmpty(t, toAccount)
	require.Equal(t, account2.ID, toAccount.ID)

	updatedAccount1, err := testRepository.GetAccount(account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testRepository.GetAccount(account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance-5, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+5, updatedAccount2.Balance)

}
