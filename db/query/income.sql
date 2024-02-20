-- name: CreateIncome :one
INSERT INTO income (
    category_id, account_id, amount, date
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: ListIncomeByAccountAndDate :many
SELECT * FROM income 
WHERE account_id = $1
AND date BETWEEN sqlc.arg(start_date) AND sqlc.arg(end_date)
ORDER BY date DESC
LIMIT $2
OFFSET $3;

-- name: GetTotalIncomeByAccountAndDate :one
SELECT SUM(amount) FROM income
WHERE account_id = $1
AND date BETWEEN sqlc.arg(start_date) AND sqlc.arg(end_date);

-- name: UpdateIncome :one
UPDATE income
SET category_id = $1, amount = $2
WHERE id = $3
RETURNING *;

-- name: DeleteIncome :exec
DELETE FROM income
WHERE id = $1;