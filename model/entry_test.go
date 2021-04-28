package model

import (
	"github.com/nebisin/gowallet/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryPayload{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testRepository.CreateEntry(arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestRepository_CreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestRepository_GetEntry(t *testing.T) {
	account := createRandomAccount(t)

	entry1 := createRandomEntry(t, account)
	entry2, err := testRepository.GetEntry(entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestRepository_ListEntry(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntryParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}

	entries, err := testRepository.ListEntry(arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, arg.AccountID)
	}
}