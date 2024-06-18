.PHONY: default
default: pyra_build pyra_run

# PYRA

.PHONY: build
build: pyra_build

.PHONY: pyra_build
pyra_build:
	@templ generate
	@go build -o ./bin/pyra ./cmd/pyra

.PHONY: pyra_build_dev
pyra_build_dev:
	@go build -tags dev -o ./tmp/bin/pyra ./cmd/pyra

.PHONY: pyra
pyra: pyra_build
	@./bin/pyra

# END PYRA

# MIGRATE

.PHONY: migrate_build
migrate_build:
	@go build -o ./bin/migrate ./cmd/migrate

.PHONY: migrate
migrate: migrate_build
	@./bin/migrate

.PHONY: rollback
rollback: migrate_build
	@./bin/migrate rollback

.PHONY: migrate_status
migrate_status: migrate_build
	@./bin/migrate status

.PHONY: migrate_version
migrate_version: migrate_build
	@./bin/migrate version

# END MIGRATE

.PHONY: seed
seed:
	@go run ./database/seeds

.PHONY: clean
clean:
	rm -rf bin/**/*

.PHONY: dev
dev:
	@air
