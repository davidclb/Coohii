-- name: GetCharacterbyID :one
SELECT * FROM character
WHERE id = @id::int LIMIT 1;


-- name: GetCharacterbyCategory :many
SELECT * FROM character
WHERE category = @category::text ;

-- name: ListCharactersByKeyset :many
SELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list
FROM character
WHERE (id > @id::int OR @id::int IS NULL)
ORDER BY id
LIMIT @pagesize::int;


-- name: FilteredCharactersByKeyset :many
SELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list
FROM character
WHERE carac_simpl ILIKE '%' || @carac::text || '%' AND (id > @id::int OR @id::int IS NULL)
ORDER BY id
LIMIT @pagesize::int;

