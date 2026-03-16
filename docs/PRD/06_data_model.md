# Modelo de Datos (Alto Nivel)

| Entidad | Campos Principales | Relaciones |
|---|---|---|
| **Usuario** | id, nombre, email, password_hash, rol | Un entrenador tiene muchos clientes |
| **Cliente** | id, usuario_id, entrenador_id, fecha_inicio | Un cliente tiene una rutina activa |
| **Rutina** | id, cliente_id, nombre, semana_inicio, activa | Una rutina tiene muchos días |
| **DiaRutina** | id, rutina_id, dia_semana (0–6) | Un día tiene muchos ejercicios |
| **Ejercicio** | id, dia_id, nombre, series, repeticiones, video_url, orden, modo_carga, peso_kg (nullable), porcentaje_carga (nullable) | Un ejercicio tiene muchos registros |
| **RegistroSesion** | id, cliente_id, fecha, tipo (autonoma/presencial), observaciones | Una sesión tiene muchos registros de ejercicio |
| **RegistroEjercicio** | id, sesion_id, ejercicio_id, series_completadas[] | Contiene peso y reps por serie |
| **PesoCorporal** | id, cliente_id, fecha, peso_kg | Pertenece a un cliente |
| **FotoProgreso** | id, cliente_id, fecha, url_foto, tipo (frontal/lateral/posterior) | Pertenece a un cliente |
| **Test1RM** | id, cliente_id, ejercicio_nombre, peso_maximo_kg, fecha, notas | Historial por ejercicio por cliente |
| **TemplateRutina** | id, entrenador_id, nombre, categoria, fecha_modificacion, veces_usado | Misma estructura de días/ejercicios que Rutina; independiente de cliente |

## Notas

- `modo_carga` en Ejercicio puede ser `fijo` (usa `peso_kg`) o `porcentaje` (usa `porcentaje_carga` + 1RM más reciente del cliente).
- `TemplateRutina` es una copia independiente; modificarla no afecta rutinas ya generadas a partir de ella.
- `RegistroEjercicio.series_completadas` almacena un arreglo de `{ peso_kg, repeticiones }` por cada serie.
