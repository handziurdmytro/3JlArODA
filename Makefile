.PHONY: gen env up down logs clean

up: env gen
	@echo "[INFO] start..."
	docker compose up --build -d
	@echo "[INFO] success!"

env:
	@echo "=== Generating .env for Crypto Service ==="
	$(MAKE) -C crypto-service gen-env

	@echo "=== Generating .env for Auth Service ==="
	@test -f auth-service/.env || ( \
		echo "PORT=3131" > auth-service/.env && \
		echo "CRYPTO_SERVICE_ADDR=crypto-service:3030" >> auth-service/.env && \
		echo "DATABASE_URL=postgres://postgres:cisco@postgres:5432/zlahoda_users?sslmode=disable" >> auth-service/.env \
	)

	@echo "=== Generating .env for Business Service ==="
	@test -f business-service/.env || ( \
		echo "PORT=2433" > business-service/.env && \
		echo "DB_URL=postgres://postgres:pswrd@postgres:5432/zlahoda_data?sslmode=disable" >> business-service/.env \
	)

	@echo "=== Generating .env for API Gateway ==="
	@test -f api-gateway/.env || ( \
		echo "PORT=8080" > api-gateway/.env && \
		echo "BUSINESS_SERVICE=zlahoda-business:2433" >> api-gateway/.env && \
		echo "AUTH_SERVICE=zlahoda-auth:3131" >> api-gateway/.env \
	)
	@echo "[INFO] all .env files are ready!"

gen:
	@echo "=== Generating Crypto Service ==="
	$(MAKE) -C crypto-service gen || true

	@echo "=== Generating Auth Service ==="
	$(MAKE) -C auth-service gen

	@echo "=== Generating Business Service ==="
	$(MAKE) -C business-service gen

	@echo "=== Generating API Gateway ==="
	$(MAKE) -C api-gateway gen
	@echo "All code generated successfully!"

down:
	docker compose down

logs:
	docker compose logs -f

clean: down
	@echo "Cleaning up environments and volumes..."
	docker compose down -v
	rm -f auth-service/.env
	rm -f business-service/.env
	rm -f api-gateway/.env
	rm -f crypto-service/.env
	@echo "Clean up finished!"

front-dbg: env gen
	@echo "Starting backend infrastructure for Frontend Debugging..."
	docker compose -f docker/test/docker-compose.front-debug.yml up --build -d
	@echo "Backend is ready! API Gateway is listening on http://localhost:8080"
	@echo "Starting frontend dev server..."
	npm --prefix frontend run dev