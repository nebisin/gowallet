postgres:
	docker run --name wallet-psql --network wallet-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:alpine

createdb:
	docker exec -it wallet-psql createdb --username=root --owner=root go_wallet

dropdb:
	docker exec -it wallet-psql dropdb go_wallet

test:
	go test -v -cover ./...

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/go_wallet?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/go_wallet?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/go_wallet?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/go_wallet?sslmode=disable" -verbose down 1

server:
	go run main.go

.PHONY: postgres createdb migrateup migratedown migrateup1 migratedown1 dropdb test server