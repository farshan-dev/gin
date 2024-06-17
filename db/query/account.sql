-- name: CreateAccount :one
INSERT INTO accounts (
  owner, blance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;