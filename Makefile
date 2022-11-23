ifneq ($(wildcard docker/.env.example),)
	ENV_FILE = .env.example
endif
ifneq ($(wildcard .env.example),)
	include .env.example
endif
ifneq ($(wildcard docker/.env),)
	ENV_FILE = .env
endif
ifneq ($(wildcard .env),)
	include .env
endif

service = backend
export

.SILENT:
.PHONY: help
help: ## Display this help screen
	awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

.PHONY: install
install: ## Installations
	go mod download
	go mod verify
	go mod tidy

.PHONY: lint
lint: ## Run linters
	golangci-lint run

.PHONY: swag-v1
swag-v1:
	swag fmt -g internal/controller/http/v1/router.go
	swag init -g internal/controller/http/v1/router.go -o docs/v1

.PHONY: swag
swag: swag-v1 ## Generate swagger docs

.PHONY: run
run: ## Run application
	go run cmd/app/main.go

.PHONY: migrate
migrate: ## Migrate database
	go run cmd/migration/main.go

.PHONY: build
build: ## Build application
	CGO_ENABLED=0 GOOS=$(OSTYPE) GOARCH=amd64 \
	go build -ldflags="-s -w" -o bin/app cmd/app/main.go

.PHONY: compose-convert
compose-convert: ## Converts the compose file to platform's canonical format
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) convert $(services)

.PHONY: compose-build
compose-build: ## Build or rebuild services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) build $(services)

.PHONY: compose-up
compose-up: ## Create and start containers
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) up -d $(services)

.PHONY: compose-down
compose-down: ## Stop and remove containers, networks
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) down --remove-orphans

.PHONY: compose-logs
compose-logs: ## View output from containers
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) logs -f $(services)

.PHONY: compose-ps
compose-ps: ## List containers
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) ps $(services)

.PHONY: compose-ls
compose-ls: ## List running compose projects
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) ls

.PHONY: compose-exec
compose-exec: ## Execute a command in a running container
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) exec $(service) sh

.PHONY: compose-start
compose-start: ## Start services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) start $(services)

.PHONY: compose-restart
compose-restart: ## Restart services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) restart $(services)

.PHONY: compose-stop
compose-stop: ## Stop services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) stop $(services)

.PHONY: docker-rm-volume
docker-rm-volume: ## Remove db volume
	docker volume rm -f flallet_db_data

.PHONY: docker-clean
docker-clean: ## Remove unused data
	docker system prune
