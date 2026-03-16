# Overview & Architecture Pattern

## What is Kochappi?

A **RESTful backend API** built in Go that powers the Kochappi platform. The API enables personal trainers to manage clients, routines, and training sessions through HTTP endpoints. It serves the Kochappi Android mobile app as the primary client, but is designed to be client-agnostic. The goal is to build a **maintainable, scalable, and testable** system that can grow from a single-trainer MVP to a multi-tenant SaaS platform.

## Core Principles

| Principle | Meaning |
|-----------|---------|
| **Independence** | Domain logic has zero dependencies on frameworks or external tools |
| **Testability** | Business logic can be tested without spinning up a database or HTTP server |
| **Flexibility** | Swap databases, auth methods, or external APIs without touching domain logic |
| **Clarity** | Clear separation of concerns — every file has one reason to exist |

---

## Hexagonal Architecture (Ports & Adapters)

We use **Hexagonal Architecture** so that the business core never depends on infrastructure.

```
┌─────────────────────────────────────────────────────────────┐
│                     HTTP ADAPTERS (API)                      │
│              (REST Handlers, Middleware, Routing)            │
└─────────┬───────────────────────────────────────────────────┘
          │
┌─────────▼───────────────────────────────────────────────────┐
│                      PORTS (Interfaces)                       │
│         (Define contracts between layers)                    │
└─────────┬───────────────────────────────────────────────────┘
          │
┌─────────▼───────────────────────────────────────────────────┐
│                  APPLICATION LAYER                           │
│          (Use Cases, Business Logic Orchestration)           │
└─────────┬───────────────────────────────────────────────────┘
          │
┌─────────▼───────────────────────────────────────────────────┐
│                   DOMAIN LAYER (Core)                        │
│        (Entities, Value Objects, Domain Rules)              │
└─────────┬───────────────────────────────────────────────────┘
          │
┌─────────▼───────────────────────────────────────────────────┐
│              ADAPTERS (Secondary/Driven)                      │
│        (PostgreSQL, File Storage, External Services)        │
└─────────────────────────────────────────────────────────────┘
```

---

## The Four Layers

### 1. Domain Layer (Core Business Logic)
- Contains entities, value objects, and domain-specific rules
- **Zero external dependencies** — pure Go, no frameworks
- Answers the question: *"what does this system do?"*
- Examples: `Trainer`, `Client`, `Routine`, `Exercise`
- Independent of HTTP, databases, or any external framework

### 2. Application Layer (Use Cases / Business Logic Orchestration)
- Orchestrates domain logic to fulfill specific business flows
- May only depend on the domain layer and ports
- Examples: `CreateRoutineUseCase`, `RegisterTrainingSessionUseCase`
- Each use case represents one piece of business functionality

### 3. Port Layer (Interfaces / Contracts)
- Defines how layers communicate without knowing about each other's implementations
- **Inbound Ports**: contracts for how external actors (HTTP clients, CLI) call into the system
- **Outbound Ports**: contracts for how the system calls external systems (database, file storage, email)
- Enables dependency inversion and easy testing

### 4. Adapter Layer (Framework Integration)
- Implements the ports — the only place where frameworks appear
- **Primary (Inbound) Adapters**: REST handlers, HTTP middleware, routing, request parsing
- **Secondary (Outbound) Adapters**: PostgreSQL/GORM repositories, JWT provider, file storage clients
- Translates between HTTP/database representations and domain objects

---

## Dependency Rule

> Dependencies always point **inward**. Inner layers never import outer layers.

```
HTTP Handler  →  Application Service  →  Domain Entity
                       ↓
               Port (interface)
                       ↑
               PostgreSQL Repository (implements port)
```
