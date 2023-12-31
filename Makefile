postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root ginko

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/ginko?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/ginko?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres dropdb ginko

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
