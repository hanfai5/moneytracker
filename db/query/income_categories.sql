-- name: CreateIncomeCategory :one
INSERT INTO income_categories (
    name, color
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetIncomeCategory :one
SELECT * FROM income_categories
WHERE id = $1 LIMIT 1;

-- name: ListIncomeCategories :many
SELECT * FROM income_categories
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateIncomeCategoryName :one
UPDATE income_categories
SET name = $1
WHERE id = $2
RETURNING *;

-- name: UpdateIncomeCategoryColor :one
UPDATE income_categories
SET color = $1
WHERE id = $2
RETURNING *;

-- name: DeleteIncomeCategory :exec
DELETE FROM income_categories
WHERE id = $1;


