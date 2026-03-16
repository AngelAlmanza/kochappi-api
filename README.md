# Kochappi Backend API — Plataforma de Gestión de Rutinas de Entrenamiento

Kochappi es un **backend API RESTful en Go** que centraliza la relación entre entrenador personal y sus clientes. Proporciona endpoints para la asignación, seguimiento y gestión integral de rutinas de entrenamiento con comunicación en tiempo real. El API sirve principalmente a la aplicación Android nativa Kochappi, pero está diseñado para ser agnóstico al cliente.

## 🎯 Características Principales (API)

- **REST Endpoints** para gestión de rutinas, clientes y entrenamientos
- **Gestión de Rutinas**: Creación y asignación de rutinas semanales personalizadas por cliente
- **Seguimiento de Progreso**: Endpoints para consultar gráficas de peso corporal, fotos y registro de ejercicios
- **Test de 1RM**: Cálculo automático de cargas por porcentaje mediante API
- **Registro de Sesiones**: Endpoints para anotaciones del entrenador durante sesiones presenciales
- **Autenticación JWT**: Diferenciación completa entre entrenadores y clientes
- **Persistencia en PostgreSQL**: Datos duraderos con migraciones versionadas
- **Arquitectura Hexagonal**: Código desacoplado, fácil de testear y mantener

## 📋 Especificaciones Técnicas

| Aspecto | Detalles |
|---|---|
| **Lenguaje** | Go 1.21+ |
| **Framework HTTP** | Gin |
| **Base de Datos** | PostgreSQL 15+ |
| **Autenticación** | JWT (RS256) |
| **Arquitectura** | Hexagonal (Ports & Adapters) |
| **API Style** | RESTful JSON |
| **Testing** | Unit + Integration + E2E |
| **Deployment** | Docker + Kubernetes ready |

## 👥 Actores del Sistema (Consumidores de la API)

### Entrenador
- Consume endpoints para gestionar múltiples clientes
- Diseña y asigna rutinas semanales vía API
- Monitorea cumplimiento y progreso consultando endpoints de reportes
- Registra observaciones de sesiones presenciales
- Accede a dashboards con métricas por cliente

### Cliente
- Consume endpoints para visualizar rutina semanal asignada
- Registra peso, repeticiones y completitud de ejercicios
- Sube fotos de progreso a través de endpoints de almacenamiento
- Participa en test de 1RM consultando endpoints de cálculo
- Consulta su evolución física mediante endpoints de reportes

## 📁 Estructura del Código

La arquitectura sigue el patrón **Hexagonal (Ports & Adapters)** para separar la lógica de negocio de la infraestructura:

```
kochappi-api/
├── cmd/
│   └── api/
│       └── main.go                 # Punto de entrada de la aplicación
│
├── internal/                        # Código privado (no exportable)
│   ├── domain/                     # Capa de dominio (lógica de negocio pura)
│   │   ├── entity/                 # Entidades core (Trainer, Client, Routine, Exercise, etc)
│   │   ├── value_object/           # Objetos de valor (Email, Weight, Password, etc)
│   │   └── error/                  # Errores específicos del dominio
│   │
│   ├── application/                # Capa de aplicación (orquestación de casos de uso)
│   │   ├── service/                # Use cases (CreateRoutine, RegisterSession, etc)
│   │   ├── dto/                    # Data Transfer Objects (request/response shapes)
│   │   └── port/                   # Puertos (interfaces de contrato)
│   │
│   ├── adapter/                    # Capa de adaptadores (implementaciones concretas)
│   │   ├── http/                   # Adaptador HTTP/REST (handlers, middleware, router)
│   │   ├── persistence/            # Adaptador de persistencia
│   │   │   ├── postgres/           # Implementación PostgreSQL + GORM
│   │   │   │   ├── migrations/     # SQL migrations versionadas
│   │   │   │   └── model/          # Modelos de base de datos
│   │   │   └── mock/               # Mocks para testing
│   │   ├── auth/                   # Adaptador de autenticación (JWT, password hashing)
│   │   └── config/                 # Configuración de la app
│   │
│   └── shared/                     # Utilidades compartidas (logger, validator, pagination)
│
├── test/                            # Tests de integración y E2E
│   ├── integration/                # Tests contra BD real
│   ├── fixtures/                   # Datos de prueba
│   └── docker-compose.test.yml     # BD para tests
│
├── config/
│   └── .env.example                # Variables de entorno
│
├── scripts/
│   ├── migrate.sh                  # Script de migraciones
│   └── seed.sh                     # Script para poblar datos
│
├── Dockerfile                      # Compilación multietapa
├── docker-compose.yml              # Servicios locales (PostgreSQL + API)
├── Makefile                        # Comandos de desarrollo
├── go.mod & go.sum                 # Dependencias
│
├── docs/                           # Documentación (ver sección Documentación abajo)
│   ├── PRD/                        # Product Requirements
│   └── architecture/               # Especificaciones técnicas
│
├── CLAUDE.md                       # Guía de desarrollo
└── README.md                       # Este archivo
```

**Ver más detalles en** [`docs/architecture/02_project_structure.md`](docs/architecture/02_project_structure.md)

## 📚 Documentación

### Producto
Accede a la documentación del producto en [`docs/PRD/`](docs/PRD/README.md):
- **Overview**: Resumen ejecutivo y objetivos
- **Scope & Roles**: Alcance del producto y definición de usuarios
- **Requisitos Funcionales**: 52 requerimientos detallados (RF-01 a RF-52)
- **Requisitos No Funcionales**: Rendimiento, offline, seguridad y usabilidad
- **Flujos**: Flujos principales de usuario
- **Modelo de Datos**: Estructura de datos de alto nivel
- **Métricas y Riesgos**: Criterios de éxito y supuestos

### Arquitectura
Consulta la arquitectura técnica en [`docs/architecture/`](docs/architecture/README.md):
- **Estructura del Proyecto**: Organización del código
- **Conceptos Core**: Patrones y decisiones arquitectónicas
- **Convenciones**: Reglas de codificación
- **Estrategia de Testing**: Enfoque de tests
- **Dependencias**: Herramientas y librerías
- **Workflow de Desarrollo**: Flujo de trabajo
- **Deployment**: Estrategia de despliegue

## 🚀 Getting Started

### Requisitos Previos
- Go 1.21+
- PostgreSQL 15+ (o Docker)
- Git
- Make (opcional, para usar Makefile)

### Instalación Rápida
```bash
# Clonar repositorio
git clone https://github.com/AngelAlmanza/kochappi-api.git
cd kochappi-api

# Copiar archivo de configuración
cp config/.env.example config/.env.local

# Instalar dependencias
go mod download

# Iniciar PostgreSQL con Docker
docker-compose up -d

# Ejecutar migraciones
make migrate

# Iniciar servidor de desarrollo con hot reload
make dev
```

El servidor estará disponible en `http://localhost:8080`

Para más detalles, ver [docs/architecture/07_development_workflow.md](docs/architecture/07_development_workflow.md)

## 📝 Convenciones de Desarrollo

Este proyecto sigue las convenciones documentadas en:
- [`CLAUDE.md`](CLAUDE.md) — Guía de desarrollo
- [`docs/architecture/04_rules_conventions.md`](docs/architecture/04_rules_conventions.md) — Convenciones técnicas

**Commit Format**: `type(scope): description` — ej: `feat(auth): add login endpoint`

