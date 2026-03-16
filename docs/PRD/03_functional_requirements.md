# Requerimientos Funcionales

## RF-01–04 · Autenticación y Gestión de Cuentas

| ID | Requisito |
|---|---|
| RF-01 | El entrenador puede crear una cuenta con correo y contraseña |
| RF-02 | El entrenador puede registrar clientes; estos reciben un acceso (invitación o credenciales generadas) |
| RF-03 | El cliente inicia sesión con sus credenciales y ve únicamente su información |
| RF-04 | El sistema soporta recuperación de contraseña vía correo electrónico |

---

## RF-05–10 · Gestión de Rutinas (Entrenador)

| ID | Requisito |
|---|---|
| RF-05 | El entrenador puede crear una rutina semanal asignada a un cliente específico |
| RF-06 | Cada rutina contempla los 7 días; cada día puede tener cero o más ejercicios |
| RF-08 | El entrenador puede editar una rutina existente; los cambios aplican desde la semana siguiente o de inmediato |
| RF-09 | El entrenador puede duplicar una rutina de un cliente para asignarla a otro |
| RF-10 | Una rutina puede marcarse como activa o inactiva |

**RF-07 — Datos por ejercicio:**
- Nombre del ejercicio
- Número de series y repeticiones objetivo
- Enlace a video externo (YouTube u otra URL) — opcional
- Porcentaje de carga (% del 1RM) — opcional; si se define, el cliente ve el peso calculado automáticamente
- Orden de ejecución dentro del día

---

## RF-11–14 · Visualización de Rutina (Cliente)

| ID | Requisito |
|---|---|
| RF-11 | El cliente ve un calendario semanal indicando los días con ejercicios asignados |
| RF-12 | Al seleccionar un día, el cliente ve la lista de ejercicios en el orden recomendado |
| RF-13 | Cada ejercicio muestra: nombre, series, repeticiones objetivo y botón de video si existe |
| RF-14 | El cliente puede navegar entre semanas para ver rutinas pasadas o futuras |

---

## RF-15–18 · Registro de Entrenamiento (Cliente)

| ID | Requisito |
|---|---|
| RF-15 | El cliente puede registrar por cada serie: peso utilizado y repeticiones realizadas |
| RF-16 | El cliente puede marcar un ejercicio o serie como completado |
| RF-17 | El cliente puede añadir una nota opcional por ejercicio (ej. lesión, fatiga) |
| RF-18 | El sistema guarda automáticamente sin requerir acción explícita de guardado |

---

## RF-19–22 · Seguimiento de Progreso (Cliente)

| ID | Requisito |
|---|---|
| RF-19 | El cliente puede registrar su peso corporal una vez por semana |
| RF-20 | El cliente puede subir fotos de progreso (frontal, lateral, posterior) |
| RF-21 | El cliente puede ver una galería cronológica de sus fotos de progreso |
| RF-22 | El cliente puede ver una gráfica de su peso corporal a lo largo del tiempo |

---

## RF-23–28 · Dashboard del Entrenador

| ID | Requisito |
|---|---|
| RF-23 | El entrenador puede seleccionar un cliente y ver su historial de entrenamientos |
| RF-24 | Por sesión: ejercicios realizados, series completadas vs planeadas, peso por serie |
| RF-25 | Gráficas por ejercicio mostrando la evolución del peso levantado en el tiempo |
| RF-26 | Indicador visual de si el cliente aumentó, mantuvo o disminuyó la carga respecto a la sesión anterior |
| RF-27 | Visualización de cumplimiento diario: 100%, parcial o no realizado |
| RF-28 | Evolución del peso corporal y fotos de progreso del cliente |

---

## RF-29–33 · Sesiones Presenciales (Entrenador)

| ID | Requisito |
|---|---|
| RF-29 | El entrenador puede registrar una sesión presencial para un cliente en una fecha específica |
| RF-30 | En sesión presencial, el entrenador registra los mismos datos que registraría el cliente |
| RF-31 | El entrenador puede añadir observaciones de texto libre por sesión |
| RF-32 | Las sesiones presenciales quedan marcadas con un indicador visual distinto en el historial |
| RF-33 | El cliente puede ver las observaciones del entrenador en sus sesiones |

---

## RF-34–42 · Test de 1RM y Porcentaje de Carga

| ID | Requisito |
|---|---|
| RF-34 | El entrenador puede registrar un test de 1RM para un cliente, asociado a un ejercicio y fecha |
| RF-35 | El sistema almacena el historial de tests de 1RM por ejercicio |
| RF-37 | Cuando un ejercicio usa % de carga, la app calcula el peso objetivo usando el 1RM más reciente (ej. "75% → 82.5 kg") |
| RF-38 | Si se asigna % de carga sin 1RM registrado, la app muestra una advertencia |
| RF-39 | El cliente puede ver el 1RM de referencia y su fecha en el detalle del ejercicio |
| RF-40 | El entrenador puede registrar un nuevo test de 1RM desde el perfil del cliente o desde una sesión presencial |
| RF-41 | Al actualizar el 1RM, las rutinas activas recalculan el peso objetivo; los registros históricos no se modifican |
| RF-42 | El entrenador puede ver una gráfica de evolución del 1RM por ejercicio en el dashboard del cliente |

**RF-36 — Modos de carga al crear un ejercicio:**
- **Peso fijo:** el entrenador especifica directamente el peso en kg
- **Porcentaje de carga:** el entrenador define un % del 1RM (ej. 75%)

---

## RF-43–52 · Templates de Rutina (Entrenador)

| ID | Requisito |
|---|---|
| RF-43 | El entrenador puede crear templates de rutina semanal independientes de cualquier cliente |
| RF-44 | Un template tiene la misma estructura que una rutina (días, ejercicios, series, reps, carga, video, orden) |
| RF-45 | Los templates son privados del entrenador; los clientes no tienen acceso |
| RF-46 | El entrenador puede organizar templates con nombre descriptivo y categoría libre (ej. "Full Body", "Pierna avanzado") |
| RF-47 | Al crear una rutina para un cliente, el entrenador puede seleccionar un template como punto de partida. El sistema crea una copia independiente |
| RF-48 | Una vez copiado el template, el entrenador puede modificar cualquier campo libremente |
| RF-49 | El entrenador puede guardar una rutina existente de un cliente como nuevo template |
| RF-50 | El entrenador puede editar, renombrar y eliminar templates. Eliminar un template no afecta rutinas ya creadas |
| RF-51 | El entrenador puede duplicar un template para crear variantes (ej. "Pecho - Volumen" y "Pecho - Fuerza") |
| RF-52 | La pantalla de templates muestra: nombre, categoría, días con ejercicios, fecha de última modificación y veces usado |

---

## Priorización

| Funcionalidad | Prioridad | Rol |
|---|---|---|
| Autenticación y roles | 🔴 Alta | Ambos |
| Creación de rutinas | 🔴 Alta | Entrenador |
| Visualización de rutina en calendario | 🔴 Alta | Cliente |
| Registro de entrenamiento | 🔴 Alta | Cliente |
| Dashboard de progreso | 🔴 Alta | Entrenador |
| Test de 1RM | 🔴 Alta | Entrenador |
| Porcentaje de carga en rutina | 🔴 Alta | Entrenador |
| Templates de rutina | 🔴 Alta | Entrenador |
| Sesiones presenciales | 🟡 Media | Entrenador |
| Peso corporal semanal | 🟡 Media | Cliente |
| Fotos de progreso | 🟡 Media | Cliente |
| Gráfica evolución 1RM | 🟡 Media | Entrenador |
| Guardar rutina como template | 🟡 Media | Entrenador |
| Notas del cliente por ejercicio | 🟢 Baja | Cliente |
| Duplicar rutinas entre clientes | 🟢 Baja | Entrenador |
