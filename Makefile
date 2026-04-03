.PHONY: build up down test

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

test:
	go test ./...

migrate-up:
	docker-compose exec business-service migrate -path /app/internal/database/migrations -database $(DATABASE_URL) up

migrate-down:
	docker-compose exec business-service migrate -path /app/internal/database/migrations -database $(DATABASE_URL) down
