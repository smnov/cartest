BIN_DIR := bin
OUTPUT := $(BIN_DIR)/carbuild
POSTGRES_USER := root
POSTGRES_PASSWORD := password

build:
	go build -o $(OUTPUT) .

run:
	go build -o $(OUTPUT) .
	./bin/carbuild

createdb:
	docker exec -it cardb psql -U $(POSTGRES_USER) -c "CREATE DATABASE cardb;"

postgresinit:
	docker run --name cardb -p 5433 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:15-alpine

postgres:
	docker exec -it cardb psql

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/cardb?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migrations -database "postgresql://root:password@localhost:5433/cardb?sslmode=disable" -verbose down

