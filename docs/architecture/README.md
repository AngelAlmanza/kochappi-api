# Kochappi Backend API Architecture

**Stack:** Go, PostgreSQL, GORM, JWT | **Pattern:** Hexagonal Architecture

This directory contains the backend API architecture documentation split by topic for easier navigation. The API serves both the Android mobile client and potential web clients with REST endpoints.

---

## Index

| File | Description |
|------|-------------|
| [01_overview.md](./01_overview.md) | Core principles and hexagonal architecture diagram |
| [02_project_structure.md](./02_project_structure.md) | Directory layout and what lives where |
| [03_core_concepts.md](./03_core_concepts.md) | Entities, Value Objects, Repositories, Use Cases, HTTP Handlers |
| [04_rules_conventions.md](./04_rules_conventions.md) | Naming, DI, error handling, auth, and all team rules |
| [05_testing_strategy.md](./05_testing_strategy.md) | Test pyramid, unit/integration/E2E examples, test helpers |
| [06_dependencies_tools.md](./06_dependencies_tools.md) | go.mod dependencies and recommended dev tools |
| [07_development_workflow.md](./07_development_workflow.md) | Local setup, Makefile commands, PR checklist |
| [08_deployment.md](./08_deployment.md) | Docker, env vars, migrations, deployment platforms |

---

## Quick Reference: Adding a New Feature

1. **Define the entity** → `internal/domain/entity/new_feature.go`
2. **Define the port** → `internal/application/port/output_port.go`
3. **Implement repository adapter** → `internal/adapter/persistence/postgres/new_feature_repository.go`
4. **Create use case** → `internal/application/service/new_feature/create_feature.go`
5. **Create HTTP handler** → `internal/adapter/http/handler/new_feature_handler.go`
6. **Write tests** → `internal/application/service/new_feature/create_feature_test.go`
7. **Add route** → `internal/adapter/http/router.go`

## Quick Reference: Adding a New Database Table

1. Create migration → `internal/adapter/persistence/postgres/migrations/XXX_create_new_table.sql`
2. Create model → `internal/adapter/persistence/postgres/model/new_table_model.go`
3. Create domain entity → `internal/domain/entity/new_entity.go`
4. Create repository → `internal/adapter/persistence/postgres/new_entity_repository.go`
5. Add port contract → `internal/application/port/output_port.go`
