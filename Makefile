include .env
export

OUTPUT := carbuild

start: postgresinit createdb migrateup run

build:
	go build -o $(OUTPUT) .

run: build
	./$(OUTPUT)

createdb:
	docker exec -it cardb psql -U $(USERNAME) -c "CREATE DATABASE $(DB_NAME);"

postgresinit:
	docker run --name cardb -p $(DB_PORT):5432 -e POSTGRES_USER=$(USERNAME) -e POSTGRES_PASSWORD=$(PASSWORD) -d postgres:15-alpine

postgres:
	docker exec -it cardb psql

migrateup:
	migrate -path db/migrations -database "postgresql://$(USERNAME):$(PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://$(USERNAME):$(PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: run build start createdb postgresinit postgres migrateup migratedown

