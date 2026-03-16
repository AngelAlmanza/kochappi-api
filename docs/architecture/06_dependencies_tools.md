# Dependencies & Tools

---

## Core Dependencies (go.mod)

```go
module kochappi

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1          // HTTP router and framework
    gorm.io/gorm v1.25.4                      // ORM
    gorm.io/driver/postgres v1.5.4            // PostgreSQL driver for GORM
    github.com/golang-jwt/jwt/v5 v5.0.0       // JWT authentication
    github.com/google/uuid v1.4.0             // UUID generation
    golang.org/x/crypto v0.15.0              // Password hashing (bcrypt)
)
```

### Installing dependencies

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/google/uuid
go get -u golang.org/x/crypto
```

---

## Recommended Dev Tools

| Tool | Purpose | Install |
|------|---------|---------|
| `air` | Hot reload — restarts the server on file save | `go install github.com/cosmtrek/air@latest` |
| `golangci-lint` | Linter — catches style and correctness issues | `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` |
| `migrate` | Database migrations CLI | `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` |
| `sqlc` | Generates type-safe Go code from SQL queries | `go install github.com/kyleconroy/sqlc/cmd/sqlc@latest` |

---

## Why These Choices?

| Decision | Reason |
|----------|--------|
| **Gin** over net/http | Routing, middleware chaining, and JSON helpers out of the box without boilerplate |
| **GORM** | Reduces SQL boilerplate for CRUD; migrations are managed separately with `migrate` tool |
| **golang-jwt** | Well-maintained, follows RFC 7519, v5 API is clean |
| **uuid** | Deterministic, collision-resistant IDs without a DB sequence |
| **bcrypt** via crypto | Industry-standard password hashing, adaptive cost factor |
