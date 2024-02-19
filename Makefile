HOME_PATH := $(shell pwd)

DOCKER_DIR="docker-compose.yaml"
MIGRATE_SQL := $(shell cat < ./migrations/specification.sql;)
BIN := "./bin/go-template"
VERSION :=$(shell date)

build:
	go build -o=$(BIN) -ldflags="-X 'main.version=${VERSION}' -X 'github.com/AlekseyPorandaykin/go-template/cmd.homeDir=${HOME_PATH}'" .

init:
	go install golang.org/x/tools/cmd/goimports@latest

run: build
	$(BIN) -config ./configs/default.toml

run-img: build-img
	docker run $(DOCKER_IMG)

up:
	./bin/go-template web

down:
	docker-compose --file=$(DOCKER_DIR) down

recreate:
	docker-compose --file=$(DOCKER_DIR) rm -f
	docker-compose --file=$(DOCKER_DIR) pull
	docker-compose --file=$(DOCKER_DIR) up --build -d

ps:
	docker-compose --file=$(DOCKER_DIR) ps

linters:
	go vet .
	gofmt -w .
	goimports -w .
	gci write /app
	gofumpt -l -w /app
	golangci-lint run ./...
	gofmt -s -l $(git ls-files '*.go')


go-fix:
	go mod tidy
	gci write ./
	gofumpt -l -w ./

.PHONY: build run build-img run-img version test lint
