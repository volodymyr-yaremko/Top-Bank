DB_URL=postgresql://root:virtualrelay@localhost:5432/top_bank_db?sslmode=disable

postgres:

	docker run --name postgres15-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=virtualrelay -d postgres:15-alpine


createdb:

	docker exec -it postgres15-alpine createdb --username=root --owner=root top_bank_db

dropdb:

	docker exec -it postgres15-alpine dropdb top_bank_db


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

test:
	go test -v -cover ./...

.PHONY : postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test