#!make
include .env

DB_URL=postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
PGDB_VOLUME=pgdb-data
NETWORK=go-rss-network
CONTAINER=pgcont16

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=$(DB_URL)
export GOOSE_MIGRATION_DIR=sql/schema

test:
	set -a && source .env && set +a && go test -v -cover ./...

initdb: pg_volume_network pg_container createdb

pg_volume_network:
	docker volume create $(PGDB_VOLUME) && docker network create -d bridge $(NETWORK)

pg_container:
	docker run --name $(CONTAINER) -v $(PGDB_VOLUME):/var/lib/postgresql/data --network $(NETWORK) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) -d postgres:16.1-alpine

createdb:
	docker exec -it $(CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(CONTAINER) dropdb "$(DB_NAME)"

dbup:
	goose -v up

dbup1:
	goose -v up up-by-one

dbdown:
	goose -v down

dbdownto:
	goose -v down $(name)

# make new_migration name=ANyString
new_migration:
	goose create $(name) sql

sqlc:
	sqlc generate

.PHONY: test initdb pg_volume_network pg_container createdb dropdb dbup dbdown dbup1 dbdownto new_migration sqlc
