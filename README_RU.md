# Микросервис для работы с балансом пользователей

## Документация

1. Postman: [docs/postman.json](docs/postman.json)
2. Swagger: [localhost:3000/api/v1/docs](http://localhost:3000/api/v1/docs)

---

## Требования

1. Go 1.19+: [go.dev/dl](https://go.dev/dl)
2. Docker: [docs.docker.com/get-docker](https://docs.docker.com/get-docker)
3. Docker Compose: [docs.docker.com/compose/install](https://docs.docker.com/compose/install)
4. GolangCI: [golangci-lint.run/usage/install](https://golangci-lint.run/usage/install)

---

## Makefile

> Отображает все цели
```bash
make help
```

```
Usage:
  make <target>
  help                Display this help screen
  install             Installations
  lint                Run linters
  swag                Generate swagger docs
  run                 Run application
  migrate             Migrate database
  build               Build application
  compose-convert     Converts the compose file to platform's canonical format
  compose-build       Build or rebuild services
  compose-up          Create and start containers
  compose-down        Stop and remove containers, networks
  compose-logs        View output from containers
  compose-ps          List containers
  compose-ls          List running compose projects
  compose-exec        Execute a command in a running container
  compose-start       Start services
  compose-restart     Restart services
  compose-stop        Stop services
  docker-rm-volume    Remove db volume
  docker-clean        Remove unused data
```

---

## Запуск микросервиса

> Установки
```bash
make install
```

> Поднимает базу данных
```bash
make compose-up services="db"
```

> Мигрирует базу данных
```bash
make migrate
```

> Запускает приложение
```bash
make run
```

---

## Запуск микросервиса в Docker

> Создает и запускает контейнеры
```bash
make compose-up
```

---

## Переменные окружения

> Отображает все примеры переменных
```bash
cat .env.example
```

```
# Backend
BACKEND_HOST         = localhost
BACKEND_PORT         = 3000
BACKEND_CORS_ORIGINS = http://localhost, https://localhost, https://localhost:8080, http://localhost:8080

# Database
DB_HOST = localhost
DB_PORT = 5432
DB_USER = postgres
DB_PASS = postgres
DB_NAME = db

# S3
S3_REGION            = us-east-1
S3_BUCKET            = flallet
S3_ENDPOINT          = https://flaiers.s3.amazonaws.com
S3_ACCESS_KEY        = ...
S3_SECRET_ACCESS_KEY = ...
```

---

## Переменные окружения в Docker

> Отображает все примеры переменных
```bash
cat docker/.env.example
```

```
# Backend
BACKEND_HOST         = 0.0.0.0
BACKEND_PORT         = 3000
HOST_BACKEND_PORT    = 3000
BACKEND_CORS_ORIGINS = http://localhost, https://localhost, https://localhost:8080, http://localhost:8080

# Database
DB_HOST      = db
DB_PORT      = 5432
HOST_DB_PORT = 5432
DB_USER      = postgres
DB_PASS      = postgres
DB_NAME      = db

# S3
S3_REGION            = us-east-1
S3_BUCKET            = flallet
S3_ENDPOINT          = https://flaiers.s3.amazonaws.com
S3_ACCESS_KEY        = ...
S3_SECRET_ACCESS_KEY = ...

# Docker
COMPOSE_PROJECT_NAME = flallet
```
