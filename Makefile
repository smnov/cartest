include .env
export

OUTPUT := carbuild

createdb:
	psql -U $(USERNAME) -c "CREATE DATABASE $(DB_NAME);"

build:
	go build -o $(OUTPUT) .

run: build
	./$(OUTPUT)

migrateup:
	migrate -path db/migrations -database "postgresql://$(USERNAME):$(PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://$(USERNAME):$(PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: run build start createdb migrateup migratedown
