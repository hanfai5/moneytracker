-- name: CreateExpenseCategory :one
INSERT INTO expense_categories (
    name, color
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetExpenseCategory :one
SELECT * FROM expense_categories
WHERE id = $1 LIMIT 1;

-- name: ListExpenseCategories :many 
SELECT * FROM expense_categories
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateExpenseCategoryName :one
UPDATE expense_categories
SET name = $1
WHERE id = $2
RETURNING *;

-- name: UpdateExpenseCategoryColor :one
UPDATE expense_categories
SET color = $1
WHERE id = $2
RETURNING *;

-- name: DeleteExpenseCategory :exec
DELETE FROM expense_categories
WHERE id = $1;