-- name: CreateEndpoint :one
INSERT INTO endpoints (
    user_id,
    name,
    token
) VALUES (
    $1, 
    $2, 
    $3
)
RETURNING *;

-- name: FindEndpointByUserID :many
SELECT * FROM endpoints
WHERE user_id = $1;

-- name: FindEndpointByToken :one
SELECT * FROM endpoints
WHERE token = $1;

-- name: DeleteEndpointByToken :exec
DELETE FROM endpoints
WHERE token = $1
  AND user_id = $2;