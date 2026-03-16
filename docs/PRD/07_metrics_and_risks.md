# Métricas, Riesgos y Supuestos

## Métricas de Éxito (3 meses post-lanzamiento)

| Métrica | Objetivo | Cómo se mide |
|---|---|---|
| Tasa de registro diario del cliente | ≥ 70% de los días activos | Sesiones registradas / sesiones planeadas |
| Rutinas completadas por semana | ≥ 60% de clientes completan su rutina | Días al 100% / días planeados |
| Retención del entrenador | ≥ 80% activos tras 30 días | Usuarios activos mes 2 / usuarios registrados |
| Tiempo de carga de rutina | < 2 s en el 95% de los casos | Telemetría de la app |
| Crashes por sesión | < 0.5% de sesiones | Firebase Crashlytics o equivalente |
| Tests de 1RM registrados | ≥ 1 test por ejercicio principal en los primeros 30 días | Registros en Test1RM / clientes activos |

---

## Riesgos

| Riesgo | Probabilidad | Impacto | Mitigación |
|---|---|---|---|
| Baja adopción del cliente para registrar entrenamientos | Media | 🔴 Alto | UX simplificada, notificaciones push, flujo de registro en < 3 pasos |
| Problemas de sincronización offline | Media | 🔴 Alto | Persistencia local con Room DB, cola de sincronización con retry automático |
| Escalabilidad del almacenamiento de fotos | Baja | 🟡 Medio | Object storage externo desde el inicio |
| Fragmentación de Android | Media | 🟡 Medio | Testing en API 26+, uso de Jetpack para compatibilidad |
| Abandono del entrenador por setup complejo | Media | 🔴 Alto | Onboarding guiado, templates de rutina predefinidas |

---

## Supuestos

- Cada entrenador gestiona entre 1 y 50 clientes en v1.0
- El entrenador tiene acceso a internet para crear y modificar rutinas
- Los clientes cuentan con Android 8.0+ y al menos 2 GB de RAM
- Los videos son enlaces externos (YouTube, Vimeo); la app no almacena videos propios
- La app no gestiona pagos ni contratos

## Dependencias Externas

| Dependencia | Uso |
|---|---|
| Backend API REST + PostgreSQL | Datos y lógica de negocio |
| Firebase Auth o JWT propio | Autenticación |
| AWS S3 / Google Cloud Storage | Almacenamiento de fotos de progreso |
| Firebase Cloud Messaging | Notificaciones push |
| MPAndroidChart o Vico | Gráficas en la app Android |

---

## Historial de Versiones

| Versión | Fecha | Cambios |
|---|---|---|
| 1.0 | Marzo 2026 | Documento inicial — definición completa de v1.0 |
| 1.1 | Marzo 2026 | Módulo Test de 1RM y Porcentaje de Carga (RF-34 a RF-42) |
| 1.2 | Marzo 2026 | Módulo Templates de Rutina (RF-43 a RF-52) |
