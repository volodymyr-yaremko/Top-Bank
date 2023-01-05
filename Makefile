DB_URL=postgresql://root:virtualrelay@localhost:5432/top_bank_db?sslmode=disable

postgres:

	docker run --name postgres15-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=virtualrelay -d postgres:15-alpine


createdb:

	docker exec -it postgres15-alpine createdb --username=root --owner=root top_bank_db

dropdb:

	docker exec -it postgres15-alpine dropdb top_bank_db

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Ayobami-00/top-bank--Golang-Postgres-Kubernetes-gRPC-/db/sqlc Store

test:
	go test -v -cover ./...

.PHONY : postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test mock db_docs