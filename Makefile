postgres:
	docker run --name pg-cont --network bank_network -p 5432:5432 -e POSTGRES_USER=root -e \
	POSTGRES_PASSWORD=olim123 -v myvolume:/var/lib/postgresql/data -d postgres
run_dockerfile:
	docker run --name olimbank --network bank_network -p 8080:8080 -e \
	DB_SOURCE="postgresql://root:olim123@pg-cont:5432/olim_bank?sslmode=disable" olimbank:latest
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
	mockgen -package mockdb -destination db/mock/store.go github.com/alim7007/go_bank_k8s/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test mock run_dockerfile