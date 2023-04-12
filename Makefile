postgres:
	docker run --name pg-cont -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=olim123 -d postgres
createdb:
	docker exec -it pg-cont createdb --username=root --owner=root olim_bank
dropdb:
	docker exec -it pg-cont dropdb olim_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:olim123@localhost:5432/olim_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:olim123@localhost:5432/olim_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:olim123@localhost:5432/olim_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:olim123@localhost:5432/olim_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/alim7007/go_backend_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mock