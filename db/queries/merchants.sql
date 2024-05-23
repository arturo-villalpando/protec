-- name: CreateMerchant :one
INSERT INTO merchants ("merchant_name", "commission") 
VALUES ($1, $2) RETURNING *;

-- name: ReadMerchant :one
SELECT * FROM merchants WHERE "merchant_id" = $1 LIMIT 1;

-- name: UpdateMerchant :one
UPDATE merchants SET "merchant_name" = $1, "commission" = $2, "updated_at" = CURRENT_TIMESTAMP 
WHERE "merchant_id" = $3 RETURNING *;

--
-- Paginate - Control
--
-- name: PaginateCountMerchants :one
SELECT count(*) FROM merchants WHERE "merchant_name" LIKE $1;

-- name: PaginateAscMerchants :many
SELECT * FROM merchants WHERE "merchant_name" LIKE $1 ORDER BY $2 ASC LIMIT $3 OFFSET $4;

-- name: PaginateDescMerchants :many
SELECT * FROM merchants WHERE "merchant_name" LIKE $1 ORDER BY $2 DESC LIMIT $3 OFFSET $4;
