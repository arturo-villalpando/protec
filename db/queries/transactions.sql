-- name: CreateTransaction :one
INSERT INTO transactions ("merchant_id", "amount", "commission", "fee") 
VALUES ($1, $2, $3, $4) RETURNING *;
