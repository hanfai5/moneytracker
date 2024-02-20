package db

import (
	"context"
	"database/sql"
	"moneytracker/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomIncomeCategory(t *testing.T) IncomeCategories {
	arg := CreateIncomeCategoryParams{
		Name:  util.RandomName(),
		Color: util.RandomColorHexString(),
	}

	category, err := testQueries.CreateIncomeCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.Color, category.Color)

	require.NotZero(t, category.ID)
	return category
}

func TestCreateIncomeCategory(t *testing.T) {
	createRandomIncomeCategory(t)
}

func TestGetIncomeCategory(t *testing.T) {
	category1 := createRandomIncomeCategory(t)
	category2, err := testQueries.GetIncomeCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Color, category2.Color)
}

func TestListIncomeCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIncomeCategory(t)
	}

	arg := ListIncomeCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	categories, err := testQueries.ListIncomeCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestUpdateIncomeCategoryColor(t *testing.T) {
	category1 := createRandomIncomeCategory(t)

	arg := UpdateIncomeCategoryColorParams{
		Color: util.RandomColorHexString(),
		ID:    category1.ID,
	}

	category2, err := testQueries.UpdateIncomeCategoryColor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, arg.Color, category2.Color)
}

func TestUpdateIncomeCategoryName(t *testing.T) {
	category1 := createRandomIncomeCategory(t)

	arg := UpdateIncomeCategoryNameParams{
		Name: util.RandomName(),
		ID:   category1.ID,
	}

	category2, err := testQueries.UpdateIncomeCategoryName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
	require.Equal(t, category1.Color, category2.Color)
}

func TestDeleteIncomeCategory(t *testing.T) {
	category1 := createRandomIncomeCategory(t)

	err := testQueries.DeleteIncomeCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetIncomeCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}
