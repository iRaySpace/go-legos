-- name: GetEmployee :one
SELECT * FROM employees
WHERE id = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY last_name, first_name;

-- name: CreateEmployee :one
INSERT INTO employees (
    first_name,
    last_name
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: UpdateEmployee :exec
UPDATE employees
SET
    first_name = $2,
    last_name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = $1;