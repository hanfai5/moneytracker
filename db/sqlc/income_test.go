package db

import (
	"context"
	"database/sql"
	"moneytracker/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomIncome(t *testing.T) Income {
	arg := CreateIncomeParams{
		CategoryID: sql.NullInt32{Int32: util.RandomId(), Valid: true},
		AccountID:  sql.NullInt32{Int32: util.RandomId(), Valid: true},
		Amount:     float32(util.RandomAmount()),
		Date:       util.RandomDate(),
	}

	income, err := testQueries.CreateIncome(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, income)

	require.Equal(t, arg.CategoryID, income.CategoryID)
	require.Equal(t, arg.AccountID, income.AccountID)
	require.Equal(t, arg.Amount, income.Amount)
	require.WithinDuration(t, arg.Date, income.Date, time.Second)
	return income
}

func TestCreateIncome(t *testing.T) {
	createRandomIncome(t)
}
