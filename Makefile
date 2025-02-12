ifneq ("$(wildcard server/.env)","")
	include server/.env
	export $(shell sed 's/=.*//' server/.env)
endif

export PATH := $(PWD)/client/node_modules/.bin:$(PATH)

.PHONT: test
test:
	@cd server && go test -cover ./...

.PHONT: test-verbose
test-verbose:
	@cd server && go test -v -cover ./...

.PHONY: dev
dev:
	@$(MAKE) -j2 dev-server dev-client

.PHONY: build
build:
	@$(MAKE) -j2 build-server build-client

.PHONY: start
start:
	@$(MAKE) -j2 start-server start-client

.PHONY: start-server
start-server:
	./server/bin/main

.PHONY: dev-server
dev-server:
	@cd server && air

.PHONY: build-server
build-server:
	@cd server && go build -o=./bin/main ./cmd

.PHONY: build-client
build-client:
	@cd client && pnpm build

.PHONY: start-client
start-client:
	@cd client && pnpm start

.PHONY: dev-client
dev-client:
	@cd client && pnpm dev

.PHONY: icons
icons:
	@cd client && pnpm icons

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

.PHONY: minio
minio:
	minio server  --console-address :9001 /Users/michaelcorrigan/minio
