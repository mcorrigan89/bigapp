ifneq ("$(wildcard server/.env)","")
	include server/.env
	export $(shell sed 's/=.*//' server/.env)
endif

export PATH := $(PWD)/client/node_modules/.bin:$(PATH)

.PHONY: dev
dev:
	@$(MAKE) -j2 dev-server dev-client

.PHONY: start-server
start-server:
	./server/bin/main

.PHONY: dev-server
dev-server:
	@cd server && air

.PHONY: build-server
build-server:
	go build -o=./server/bin/main ./server/cmd

.PHONY: dev-client
dev-client:
	@cd client && pnpm dev

.PHONY: codegen
codegen:
	buf lint
	buf generate --template ./server/buf.gen.yaml api
	buf generate --template ./client/buf.gen.yaml api

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
