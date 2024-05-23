-- name: TransactionsEarns :one
SELECT SUM(fee)::double precision AS total_earns FROM transactions;

-- name: TransactionsEarnsByMerchant :one
SELECT SUM(fee)::double precision AS total_earns FROM transactions WHERE "merchant_id" = $1;