-- name: GetUserbyID :one
SELECT * FROM users
WHERE id = @id::int LIMIT 1;

-- name: GetUserbyUsername :one
SELECT * FROM users
WHERE username = @username::text LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  @username::text, @email::text, @password::text
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @id::int;