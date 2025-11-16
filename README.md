# 🎬 Streaming Management System 

## 🔰 Estado del Proyecto

| Fase Actual | Estado | Descripción |
| :--- | :--- | :--- |
| **Planeación (Etapa 1)** | ✅&nbsp;`COMPLETADA` | Se han definido el objetivo, módulos, estructura de directorios, entidades clave y dependencias de terceros. |
|   Implementación&nbsp;(Etapa&nbsp;2) | ⏸ Pendiente | Desarrollo del código en Go, implementando la lógica funcional y la conexión a la base de datos. |

## 📌 1. Introducción

Proyecto backend desarrollado en **Go (Golang)** con **MySQL**, siguiendo una arquitectura modular y escalable.

Este documento presenta la arquitectura, módulos, funcionalidades, tecnologías y alcance del proyecto para construir un sistema similar a Netflix o Disney+.

---
## 🚀 2. Objetivo del Proyecto

Diseñar y desarrollar un backend robusto para una plataforma de streaming que permita:

- Gestión completa de usuarios y autenticación.
- Administración de contenido multimedia (películas, series, documentales).
- Selección de planes, pagos y control de suscripciones.
- Reproducción con historial individual por perfil.
- Integración con métodos de pago (PayPal, tarjetas).

---
## 📌 3. Alcance del Sistema

### ✔ Incluye
- Backend en Go con arquitectura modular.
- Base de datos MySQL.
- Gestión de usuarios, perfiles y credenciales.
- Gestión de contenido y catálogo.
- Manejo de suscripciones y pasarelas de pago.
- Registro del historial de reproducción por perfil.
- API REST para consumo desde web o móvil.

### ❌ No Incluye (por ahora)
- Frontend web o aplicación móvil.
- Sistema de streaming o transcodificación real.
- CDN o distribución global de video.
- Recomendaciones avanzadas.

---
## 🧩 4. Arquitectura del Sistema

El sistema se divide en cuatro módulos principales:

### 3.1 Gestión de Usuarios y Autenticación
- Registro y login con correo + contraseña.
- Tokens JWT para sesiones.
- Recuperación de contraseña.
- Autenticación opcional MFA (código OTP).
- Gestión de múltiples perfiles por usuario.

### 3.2 Gestión de Contenido
- CRUD de películas, series y documentales.
- Metadatos: título, descripción, duración, fecha de publicación.
- Clasificación y filtros.
- Listado de catálogo disponible.

### 3.3 Gestión de Suscripciones y Pagos
- Elección de planes de suscripción.
- Integración con PayPal.
- Renovación automática.
- Control del estado de suscripción (activa, vencida, en pago).

### 3.4 Reproducción e Historial
- Registro del contenido reproducido por perfil.
- Continuar viendo (último punto registrado).
- Historial únicamente asociado al perfil que reproduce.

---
## 🏗 5. Estructura de Funcionalidades por Módulo

### 📁 Módulo: User/Auth
- Registro de usuario.
- Inicio de sesión con JWT.
- Validación de credenciales.
- Creación/edición/eliminación de perfiles.
- Recuperación de contraseña.
- Autenticación adicional (MFA opcional).

### 📁 Módulo: Content
- Alta, edición y eliminación de contenido.
- Filtro por categoría o tipo (película, serie, documental).
- Visualización de detalles de contenido.
- Listado de catálogo completo.

### 📁 Módulo: Subscriptions & Payments
- Selección de planes.
- Procesamiento de pagos.
- Integración con PayPal y Stripe.
- Renovaciones.
- Consulta del estado de la suscripción.

### 📁 Módulo: Playback & History
- Registro de reproducción.
- Historial por perfil.
- Continuar viendo.
- Indicadores de visualización reciente.

---
## 💻 6. Tecnologías Utilizadas

### Backend – Go (Golang)
Se eligió Go por:
- Rendimiento superior en servidores backend.
- Manejo eficiente de concurrencia.
- Estabilidad y facilidad de mantenimiento.
- Ideal para APIs y microservicios.

### Paquetes estándar de Go
- `net/http` – Servidor HTTP y manejo de rutas.
- `encoding/json` – Serialización/deserialización JSON.
- `database/sql` – Interacción con MySQL.
- `context`, `time`, `errors`.

### Paquetes externos
- `github.com/go-chi/chi/v5` – Router simple y eficiente.
- `gorm.io/gorm` – ORM para MySQL.
- `github.com/go-sql-driver/mysql` – Driver MySQL oficial.
- `github.com/golang-jwt/jwt/v5` – Tokens JWT.
- `github.com/spf13/viper` – Variables de entorno y configuración.
- SDK de PayPal.

---

## 🗄️ 7. Base de Datos – MySQL

MySQL es la base elegida por:
- Escalabilidad.
- Facilidad de administración.
- Amplio soporte en la comunidad.
- Integración nativa con Go.

---

## 📚 8. Diagrama de Entidades Clave (Modelado POO)

El modelado define las relaciones fundamentales que guiarán las estructuras de datos en Go:

* **Entidades Centrales:** `Usuario`, `Suscripción`, `Contenido`, `Historial`, `Transacción`.

![Diagrama de clases - Streaming](https://github.com/user-attachments/assets/8a547339-b2ba-4fce-8a68-d5ee16aae42a)


