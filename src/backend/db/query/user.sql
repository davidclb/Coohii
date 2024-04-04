-- name: GetUserbyID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserbyUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;