ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY: dev
dev:
	@$(MAKE) dev-server

.PHONY: start-server
start-server:
	./server/bin/main

.PHONY: dev-server
dev-server:
	cd server && air

.PHONY: build-server
build-server:
	go build -o=./server/bin/main ./server/cmd

.PHONY: codegen
codegen:
	buf lint
	buf generate

.PHONY: models
models:
	pg_dump --schema-only simple_auth > server/schema.sql
	sqlc generate -f server/sqlc.yaml

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir server/migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate -path=./server/migrations -database="$(POSTGRES_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path=./server/migrations -database="$(POSTGRES_URL)" down 1
