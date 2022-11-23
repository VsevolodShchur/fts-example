all: build down up

pull:
	docker-compose pull

push:
	docker-compose push

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

psql:
	docker exec -it ullr__db psql -U postgres

rotate-index:
	docker-compose exec -T manticore sh -c "indexer --rotate --all"

.PHONY: pull push build up down psql
