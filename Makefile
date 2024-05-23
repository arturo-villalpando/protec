.PHONY: clean
clean:
	rm -rf internal/db/*

.PHONY: swag
swag:
	swag init

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: migration
migration:
	migrate create -ext sql -dir db/migrations/ -seq init_mg

.PHONY: migrate_up
migrate_up:
	migrate -database postgres://postgres:olEK98s8n9Z8hSKH12f@127.0.0.1:5433/protec?sslmode=disable -path db/migrations/ up

.PHONY: migrate_down
migrate_down:
	migrate -database postgres://postgres:olEK98s8n9Z8hSKH12f@127.0.0.1:5433/protec?sslmode=disable -path db/migrations/ down

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build main.go