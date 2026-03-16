# Deployment

---

## 1. Docker Build

Multi-stage build to keep the final image small:

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/kochappi-api ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/bin/kochappi-api .
COPY --from=builder /app/internal/adapter/persistence/postgres/migrations ./migrations

EXPOSE 8080
CMD ["./kochappi-api"]
```

Build and run locally:

```bash
docker build -t kochappi-api .
docker run -p 8080:8080 --env-file config/.env.production kochappi-api
```

---

## 2. Environment Variables

```bash
# config/.env.production
DATABASE_URL=postgresql://user:pass@db-host:5432/kochappi
JWT_SECRET=your-secret-key-here
PORT=8080
ENV=production
LOG_LEVEL=info
```

> Never commit `.env.production` to version control. Inject secrets via your deployment platform's secret management.

---

## 3. Database Migrations

Always run migrations **before** deploying the new application version:

```bash
migrate -path ./migrations -database "$DATABASE_URL" up
```

To roll back the last migration:

```bash
migrate -path ./migrations -database "$DATABASE_URL" down 1
```

---

## 4. Recommended Deployment Platforms

| Platform | Best for | Notes |
|----------|---------|-------|
| **Railway** | Getting started fast | Pay-per-use, easy PostgreSQL add-on |
| **Render** | Free tier + simplicity | Good free PostgreSQL, auto-deploys from GitHub |
| **DigitalOcean App Platform** | Small teams, affordable | Simple managed DB add-on |
| **AWS EC2 + RDS** | Production scale | More control, more operational overhead |

---

## 5. docker-compose.yml (local dev)

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: kochappi
      POSTGRES_PASSWORD: password
      POSTGRES_DB: kochappi_dev
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  api:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - config/.env.local
    depends_on:
      - postgres

volumes:
  postgres_data:
```
