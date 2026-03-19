.PHONY: help dev test test-unit test-int migrate seed lint build run docker-up docker-down swagger

dev:
	air

test:
	go test ./... -v -cover

test-unit:
	go test ./internal/... -v

test-int:
	go test ./test/integration/... -v

migrate:
	migrate -path internal/adapter/persistence/postgres/migrations -database "$$DATABASE_URL" up

seed:
	go run scripts/seed.go

lint:
	golangci-lint run

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/kochappi-api ./cmd/api

run:
	./bin/kochappi-api

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

swagger:
	swag init -g cmd/api/main.go -o docs
