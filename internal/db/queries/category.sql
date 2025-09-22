-- name: CreateCategory :one
INSERT INTO
  categories (name, description)
VALUES
  ($1, $2) RETURNING *;

-- name: GetCategoryById :one
SELECT
  *
FROM
  categories
WHERE
  id = $1;

-- name: ListCategories :many
SELECT
  *
FROM
  categories
ORDER BY
  id
LIMIT
  $1
OFFSET
  $2;

-- name: UpdateCategory :one
UPDATE categories
SET
  name = $2,
  description = $3
WHERE
  id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE
  id = $1;