-- name: GetStorybyCharacter :many
SELECT * FROM stories
WHERE carac_ucs_id = $1 ;

-- name: GetStorybyUser :many
SELECT * FROM stories
WHERE user_id = $1 ;

-- name: CreateStory :one
INSERT INTO stories (
  user_id, carac_ucs_id, body, public
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteStory :exec
DELETE FROM stories
WHERE user_id = $1 and carac_ucs_id = $2;

