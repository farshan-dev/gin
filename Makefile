postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -d postgres:16-alpine

dbstart:
	docker start postgres16

dbstop:
	docker stope postgres16

dbcreate:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dbdrop:
	docker exec -t postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/farshan-dev/gin/db/sqlc Store

.PHONY: postgres dbstart dbstop dbcreate dbdrop migrateup migratedown sqlc test server mock