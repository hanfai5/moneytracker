postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root money_tracker

dropdb:
	docker exec -it postgres dropdb money_tracker

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/money_tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/money_tracker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup sqlc test