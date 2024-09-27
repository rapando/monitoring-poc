-- name: CreateRandomData :exec
INSERT INTO random_data (x,y) VALUES (?,?);

-- name: CountData :one
SELECT COUNT(id) AS qtty FROM random_data;
