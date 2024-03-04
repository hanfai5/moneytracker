// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: income_categories.sql

package db

import (
	"context"
)

const createIncomeCategory = `-- name: CreateIncomeCategory :one
INSERT INTO income_categories (
    name, color
) VALUES (
    $1, $2
)
RETURNING id, name, color
`

type CreateIncomeCategoryParams struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (q *Queries) CreateIncomeCategory(ctx context.Context, arg CreateIncomeCategoryParams) (IncomeCategories, error) {
	row := q.db.QueryRowContext(ctx, createIncomeCategory, arg.Name, arg.Color)
	var i IncomeCategories
	err := row.Scan(&i.ID, &i.Name, &i.Color)
	return i, err
}

const deleteIncomeCategory = `-- name: DeleteIncomeCategory :exec
DELETE FROM income_categories
WHERE id = $1
`

func (q *Queries) DeleteIncomeCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteIncomeCategory, id)
	return err
}

const getIncomeCategory = `-- name: GetIncomeCategory :one
SELECT id, name, color FROM income_categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetIncomeCategory(ctx context.Context, id int32) (IncomeCategories, error) {
	row := q.db.QueryRowContext(ctx, getIncomeCategory, id)
	var i IncomeCategories
	err := row.Scan(&i.ID, &i.Name, &i.Color)
	return i, err
}

const listIncomeCategories = `-- name: ListIncomeCategories :many
SELECT id, name, color FROM income_categories
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListIncomeCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListIncomeCategories(ctx context.Context, arg ListIncomeCategoriesParams) ([]IncomeCategories, error) {
	rows, err := q.db.QueryContext(ctx, listIncomeCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []IncomeCategories{}
	for rows.Next() {
		var i IncomeCategories
		if err := rows.Scan(&i.ID, &i.Name, &i.Color); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateIncomeCategoryColor = `-- name: UpdateIncomeCategoryColor :one
UPDATE income_categories
SET color = $1
WHERE id = $2
RETURNING id, name, color
`

type UpdateIncomeCategoryColorParams struct {
	Color string `json:"color"`
	ID    int32  `json:"id"`
}

func (q *Queries) UpdateIncomeCategoryColor(ctx context.Context, arg UpdateIncomeCategoryColorParams) (IncomeCategories, error) {
	row := q.db.QueryRowContext(ctx, updateIncomeCategoryColor, arg.Color, arg.ID)
	var i IncomeCategories
	err := row.Scan(&i.ID, &i.Name, &i.Color)
	return i, err
}

const updateIncomeCategoryName = `-- name: UpdateIncomeCategoryName :one
UPDATE income_categories
SET name = $1
WHERE id = $2
RETURNING id, name, color
`

type UpdateIncomeCategoryNameParams struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

func (q *Queries) UpdateIncomeCategoryName(ctx context.Context, arg UpdateIncomeCategoryNameParams) (IncomeCategories, error) {
	row := q.db.QueryRowContext(ctx, updateIncomeCategoryName, arg.Name, arg.ID)
	var i IncomeCategories
	err := row.Scan(&i.ID, &i.Name, &i.Color)
	return i, err
}
