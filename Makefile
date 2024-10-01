postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.4-alpine3.20

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -source db/sqlc/store.go -destination db/mock/store.go -package mockdb

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock
