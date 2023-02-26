package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTX(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfers from account1 to account2 (use goroutines)
	n := 5
	amount := int64(10)

	// ? Two channels are created to receive the results & Errors of the goroutines
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// ? The results are received from the channels
	// ? The results are checked for errors
	// ? The results are checked for the correct values
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// ? check transfer records
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.Equal(t, amount, result.Transfer.Amount)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		// ? Now to make sure that transfer record was really created in the database
		_, err = testQueries.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		// ? Check fromEntry records
		require.NotEmpty(t, result.FromEntry)
		require.Equal(t, account1.ID, result.FromEntry.AccountID)
		require.Equal(t, -amount, result.FromEntry.Amount)
		require.NotZero(t, result.FromEntry.ID)
		require.NotZero(t, result.FromEntry.CreatedAt)

		// ? Now to make sure that entry record was really created in the database
		_, err = testQueries.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)

		// ? Check toEntry records
		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, account2.ID, result.ToEntry.AccountID)
		require.Equal(t, amount, result.ToEntry.Amount)
		require.NotZero(t, result.ToEntry.ID)
		require.NotZero(t, result.ToEntry.CreatedAt)

		// ? Now to make sure that entry record was really created in the database
		_, err = testQueries.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)

		// TODO: check account balances
	}
}
