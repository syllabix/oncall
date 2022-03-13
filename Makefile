## Print the help message.
# Parses this Makefile and prints targets that are preceded by "##" comments.
help:
	@echo "" >&2
	@echo "Available targets: " >&2
	@echo "" >&2
	@awk -F : '\
			BEGIN { in_doc = 0; } \
			/^##/ && in_doc == 0 { \
				in_doc = 1; \
				doc_first_line = $$0; \
				sub(/^## */, "", doc_first_line); \
			} \
			$$0 !~ /^#/ && in_doc == 1 { \
				in_doc = 0; \
				if (NF <= 1) { \
					next; \
				} \
				printf "  %-15s %s\n", $$1, doc_first_line; \
			} \
			' <"$(abspath $(lastword $(MAKEFILE_LIST)))" \
		| sort >&2
	@echo "" >&2

name = ""

## run the server locally as configured by a .env file in the root of backend dir
run:
	go1.18beta2 run main.go

## run go mod tidy
tidy:
	go1.18beta2 mod tidy

## Start development environment (generates code, spins up databases, etc)
dev.start: 
	export LOCAL_MOUNT=$(shell pwd)
	mkdir -p .cache/pkg	
	docker compose up -d

## Stop the development environment
dev.stop:
	docker compose down

## Sets up Slack bot manifest generator utility
manifest:
	go1.18beta2 run ./.dev/manifest/generator.go

## Creates a new entity. (Ex: make entity name=User)
entity:
	go run entgo.io/ent/cmd/ent init $(name) --target=./datastore/entity/schema

## Updates entities
entity.update:
	go generate ./datastore/entity/...

## release to heroku
release:
	heroku container:login
	heroku container:push web
	heroku container:release web

## Creates a new db migration file for the provided service. (Ex: make migration name=cool-new-tables)
migration:
	$(MAKE) _dexec CMD="sql-migrate new -config=common/db/migrations/dbconfig.yml -env=dev $(name)"

## Run an up migration
migrate.up:
	$(MAKE) _dexec CMD="sql-migrate up -config=common/db/migrations/dbconfig.yml -env=dev"

## Run a down migration
migrate.down:
	$(MAKE) _dexec CMD="sql-migrate down -config=common/db/migrations/dbconfig.yml -env=dev"

## Generates database models based upon the existing tables in your local development database
dbmodels:
	rm -rf ./datastore/model/*.go
	$(MAKE) _dexec CMD="sqlboiler -c datastore/model/sqlboiler.toml -o datastore/model -p model --tag db --no-hooks psql --no-tests"

_dexec:
	docker exec oncall_dev ${CMD}