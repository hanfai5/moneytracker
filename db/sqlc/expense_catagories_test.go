package db

import (
	"context"
	"database/sql"
	"moneytracker/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomExpenseCategory(t *testing.T) ExpenseCategories {
	arg := CreateExpenseCategoryParams{
		Name:  util.RandomName(),
		Color: util.RandomColorHexString(),
	}

	category, err := testQueries.CreateExpenseCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.Color, category.Color)

	require.NotZero(t, category.ID)
	return category
}

func TestCreateExpenseCategory(t *testing.T) {
	createRandomExpenseCategory(t)
}

func TestGetExpenseCategory(t *testing.T) {
	category1 := createRandomExpenseCategory(t)

	category2, err := testQueries.GetExpenseCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Color, category2.Color)
}

func TestListExpenseCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomExpenseCategory(t)
	}

	arg := ListExpenseCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	categories, err := testQueries.ListExpenseCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestUpdateExpenseCategoryColor(t *testing.T) {
	category1 := createRandomExpenseCategory(t)

	arg := UpdateExpenseCategoryColorParams{
		Color: util.RandomColorHexString(),
		ID:    category1.ID,
	}

	category2, err := testQueries.UpdateExpenseCategoryColor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, arg.Color, category2.Color)
}

func TestUpdateExpenseCategoryName(t *testing.T) {
	category1 := createRandomExpenseCategory(t)

	arg := UpdateExpenseCategoryNameParams{
		Name: util.RandomName(),
		ID:   category1.ID,
	}

	category2, err := testQueries.UpdateExpenseCategoryName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
	require.Equal(t, category1.Color, category2.Color)
}

func TestDeleteExpenseCategory(t *testing.T) {
	category1 := createRandomExpenseCategory(t)

	err := testQueries.DeleteExpenseCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetExpenseCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}
