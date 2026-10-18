-- name: ListGames :many
SELECT * FROM games;


-- name: ListGamesById :one
SELECT * FROM games WHERE id = $1;
