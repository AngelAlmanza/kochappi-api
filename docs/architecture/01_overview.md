# Overview & Architecture Pattern

## What is Kochappi?

A backend API for personal trainers to manage clients, routines, and training sessions. The goal is to build a **maintainable, scalable, and testable** system that can grow from a single-trainer MVP to a multi-tenant SaaS platform.

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

### 2. Application Layer (Use Cases)
- Orchestrates domain logic to fulfill specific business flows
- May only depend on the domain layer
- Examples: `CreateRoutineUseCase`, `RegisterTrainingSessionUseCase`

### 3. Port Layer (Interfaces / Contracts)
- Defines how layers communicate without knowing about each other
- **Inbound Ports**: how external actors call into the system (HTTP → Application)
- **Outbound Ports**: how the system calls external systems (Application → Database)

### 4. Adapter Layer (Framework Integration)
- Implements the ports — the only place where frameworks appear
- **Primary (Inbound)**: REST handlers, middleware, routing
- **Secondary (Outbound)**: PostgreSQL/GORM repositories, JWT provider, file storage

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
