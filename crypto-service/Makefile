.PHONY: gen gen-env docker-build docker-run

gen:
	@echo "Crypto-service generates protos via build.rs automatically. Nothing to do."

gen-env:
	@echo  "Creating .env file..."
	@echo  "PORT=3030" > .env
	@echo  "JWT_EXP_SECONDS=86400" >> .env
	@echo  "PEPPER=$$(openssl rand -hex 32)" >> .env
	@echo  "SIGNING_KEY=$$(openssl rand -hex 32)" >> .env
	@echo  "JWT_SECRET=$$(openssl rand -hex 32)" >> .env
	@echo  "Successfully created"

docker-build:
	cd .. && \
	docker build -t zlahoda-crypto-service -f crypto-service/Dockerfile .

docker-run:
	cd .. && \
	docker run -p 3030:3030 --env-file ./crypto-service/.env zlahoda-crypto-service