.PHONY: gen-env

gen-env:
	@echo	"Creating .env file..."
	@echo	"PORT=3030" > .env
	@echo	"JWT_EXP_SECONDS=86400" >> .env
	@echo	"PEPPER=$$(openssl rand -hex 32)" >> .env
	@echo	"SIGNING_KEY=$$(openssl rand -hex 32)" >> .env
	@echo	"JWT_SECRET=$$(openssl rand -hex 32)" >> .env
	@echo	"Successfully created"