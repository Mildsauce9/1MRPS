-- name: ListGames :many
SELECT * FROM games;


-- name: ListGames :one
SELECT * FROM games WHERE id = $1;
