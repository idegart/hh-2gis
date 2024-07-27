start: build up.d

restart: down up.d

up:
	@docker compose up

build:
	@docker compose build

up.d:
	@docker compose up -d

down:
	@docker compose down

logs:
	@docker compose logs