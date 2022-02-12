NAME := app
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./cmd
DEV_CONFIG_PATH := ./configs/dev.yml
CONFIG_TEMPLATE_PATH := ./configs/template.yml

.PHONY: run
run: entgen
	go run cmd/$(NAME)/main.go ${DEV_CONFIG_PATH}

.PHONY: test
test:
	go test ./... -cover

.PHONY: build
build:
	echo "building"
	${BUILD_CMD}
	echo "build done"

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.29.0
	./bin/golangci-lint run -v

.PHONY: entgen
entgen:
	cd ./internal/repository/entgo && go generate ./ent

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d --force-recreate --build --remove-orphans --always-recreate-deps --renew-anon-volumes