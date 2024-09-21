package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mauro-moreno/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, 0)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	args := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, args.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)

	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, 0)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, acc := range accounts {
		require.NotEmpty(t, acc)
		require.NotZero(t, acc.ID)
		require.NotZero(t, acc.CreatedAt)
	}
}

func createRandomAccount(t *testing.T) Account {
	account := Account{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    account.Owner,
		Balance:  account.Balance,
		Currency: account.Currency,
	})
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, account.Owner, acc.Owner)
	require.Equal(t, account.Balance, acc.Balance)
	require.Equal(t, account.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}
