// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    $1
) RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i Users
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i Users
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name FROM users
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $1
WHERE id = $2
RETURNING id, name
`

type UpdateUserParams struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Name, arg.ID)
	var i Users
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
