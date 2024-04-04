-- name: GetCharacterbyID :one
SELECT * FROM characters
WHERE ucs_id = $1 LIMIT 1;

-- name: GetCharacterbyHeisigIndexSimpl :one
SELECT * FROM characters
WHERE index_heisig_simpl = $1 ;

-- name: GetCharacterbyHeisigIndexTrad :one
SELECT * FROM characters
WHERE index_heisig_trad = $1 ;

