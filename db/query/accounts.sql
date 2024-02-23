-- name: CreateAccount :one
INSERT INTO accounts (
    user_id, balance
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY balance
LIMIT $1
OFFSET $2;