postgres:
	docker run --name postgres-instance-1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345678 -d postgres:12-alpine

createdb:
	docker exec -it postgres-instance-1 createdb --username=root --owner=root bankitup

dropdb:
	docker exec -it postgres-instance-1 dropdb bankitup

migrationup:
	migrate -path db/migration -database "postgres://root:12345678@localhost:5432/bankitup?sslmode=disable" up

migrationdown:
	migrate -path db/migration -database "postgres://root:12345678@localhost:5432/bankitup?sslmode=disable" down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrationup migrationdown sqlc