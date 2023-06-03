DB_URL=postgresql://root:olim123@localhost:5432/olimbank?sslmode=disable

postgres:
	docker run --name postgresdb --network bank_network -p 5432:5432 -e POSTGRES_USER=root -e \
	POSTGRES_PASSWORD=olim123 -v myvolume:/var/lib/postgresql/data -d postgres
run_dockerfile:
	docker run --name olimbank --network bank_network -p 8080:8080 -e \
	DB_SOURCE="postgresql://root:olim123@postgresdb:5432/olimbank?sslmode=disable" olimbank:0.1
createdb:
	docker exec -it postgresdb createdb --username=root --owner=root olimbank
dropdb:
	docker exec -it postgresdb dropdb olimbank
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
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/alim7007/go_bank_k8s/db/sqlc Store
db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=olimbank \
	proto/*.proto
	
evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test mock run_dockerfile db_docs db_schema proto evans