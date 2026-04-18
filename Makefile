.PHONY: build up down test test-auth-crypto test-auth-crypto-down

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

test-auth-crypto:
	docker compose -f docker/test/docker-compose.auth-crypto.yml up --build -d

test-auth-crypto-down:
	docker compose -f docker/test/docker-compose.auth-crypto.yml down -v