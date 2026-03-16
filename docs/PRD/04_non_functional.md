# Requerimientos No Funcionales

## Rendimiento

- Rutina del día carga en < 2 segundos con conexión 4G estándar
- Gráficas del entrenador renderizan en < 3 segundos para historial de hasta 12 meses
- El registro de una serie se guarda localmente de inmediato; sincronización al servidor en background

## Disponibilidad y Offline

- El cliente puede registrar su entrenamiento sin internet. La sincronización ocurre al recuperar conectividad
- La rutina de la semana actual debe estar disponible offline una vez descargada

## Seguridad

- Toda comunicación app ↔ servidor va sobre HTTPS/TLS 1.2 o superior
- Las fotos de progreso son privadas: solo el cliente y su entrenador pueden acceder a ellas
- Las contraseñas se almacenan con hash (bcrypt o equivalente); nunca en texto plano
- Las sesiones inactivas expiran después de 30 días sin refrescar el token

## Usabilidad

- Navegable con una sola mano en pantallas de 5" a 6.7"
- Elementos táctiles con tamaño mínimo de 48×48 dp (Material Design)
- El flujo de registro de una serie (peso + reps + confirmar) no supera 3 interacciones

## Compatibilidad

- Android 8.0 (API 26) o superior
- Funciona correctamente desde 360 dp de ancho de pantalla
- Soporte para modo oscuro de Android (Android 10+)

## Escalabilidad

- Backend soporta al menos 500 usuarios activos simultáneos en fase inicial con arquitectura escalable horizontalmente
- Las fotos se almacenan en object storage externo (AWS S3 o equivalente); no en el servidor de aplicaciones
