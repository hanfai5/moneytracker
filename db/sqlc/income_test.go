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
	category := createRandomIncomeCategory(t)
	account := createRandomAccount(t)

	arg := CreateIncomeParams{
		CategoryID: sql.NullInt32{Int32: category.ID, Valid: true},
		AccountID:  sql.NullInt32{Int32: account.ID, Valid: true},
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

func TestGetIncome(t *testing.T) {
	income1 := createRandomIncome(t)

	income2, err := testQueries.GetIncome(context.Background(), income1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, income2)

	require.Equal(t, income1.ID, income2.ID)
	require.Equal(t, income1.CategoryID, income2.CategoryID)
	require.Equal(t, income1.AccountID, income2.AccountID)
	require.Equal(t, income1.Amount, income2.Amount)
	require.WithinDuration(t, income1.Date, income2.Date, time.Second)
}

func TestGetTotalIncomeByAccountAndDate(t *testing.T) {
	account := createRandomAccount(t)

	arg := GetTotalIncomeByAccountAndDateParams{
		AccountID: sql.NullInt32{Int32: account.ID, Valid: true},
		StartDate: util.RandomDate(),
		EndDate:   util.RandomDate(),
	}

	totalIncome, err := testQueries.GetTotalIncomeByAccountAndDate(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, totalIncome)
}

func TestListIncomeByAccountAndDate(t *testing.T) {
	account := createRandomAccount(t)

	arg := ListIncomeByAccountAndDateParams{
		AccountID: sql.NullInt32{Int32: account.ID, Valid: true},
		StartDate: util.RandomDate(),
		EndDate:   util.RandomDate(),
		Limit:     5,
		Offset:    5,
	}

	_, err := testQueries.ListIncomeByAccountAndDate(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateIncome(t *testing.T) {
	income1 := createRandomIncome(t)

	arg := UpdateIncomeParams{
		ID:         income1.ID,
		Amount:     float32(util.RandomAmount()),
		CategoryID: income1.CategoryID,
	}

	income2, err := testQueries.UpdateIncome(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, income2)
	require.Equal(t, arg.ID, income2.ID)
	require.Equal(t, arg.Amount, income2.Amount)
	require.Equal(t, arg.CategoryID, income2.CategoryID)
	require.Equal(t, income1.AccountID, income2.AccountID)
	require.WithinDuration(t, income1.Date, income2.Date, time.Second)
}

func TestDeleteIncome(t *testing.T) {
	income1 := createRandomIncome(t)

	err := testQueries.DeleteIncome(context.Background(), income1.ID)
	require.NoError(t, err)

	income2, err := testQueries.GetIncome(context.Background(), income1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, income2)
}
