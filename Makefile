.PHONY: default
default: pyra_build pyra_run

.PHONY: build
build: pyra_build

.PHONY: pyra_build
pyra_build:
	@templ generate
	@go build -o ./bin/pyra ./cmd/pyra

.PHONY: pyra_run
pyra_run: pyra_build
	@./bin/pyra

.PHONY: clean
clean:
	rm -rf bin/**/*

.PHONY: dev
dev:
	@air
