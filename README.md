# 🎬 Streaming Management System 

## 🔰 Estado del Proyecto

| Fase Actual | Estado | Descripción |
| :--- | :--- | :--- |
| **Planeación (Etapa 1)** | ✅&nbsp;`COMPLETADA` | Se han definido el objetivo, módulos, estructura de directorios, entidades clave y dependencias de terceros. |
|   Implementación&nbsp;(Etapa&nbsp;2) | ⏸ Proceso | Desarrollo del código en Go, implementando la lógica funcional y la conexión a la base de datos. |

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

## 📚 8. Diagrama de Clases

![Diagrama de clases - Streaming](https://github.com/user-attachments/assets/8a547339-b2ba-4fce-8a68-d5ee16aae42a)


## 🚀 Características Principales

### 👤 Gestión de Usuarios y Perfiles
* **Autenticación Completa:** Registro e inicio de sesión seguro para usuarios.
* **Selección de Perfil:** Pantalla intermedia estilo "Netflix" que permite elegir o crear perfiles personalizados después del login.
* **Control de Acceso:** Sistema que verifica suscripciones activas antes de permitir el acceso al catálogo.

### 🎬 Experiencia del Usuario (Dashboard)
* **Visualización Intuitiva:** Catálogo organizado en una grilla moderna con títulos y descripciones siempre visibles para mejorar la navegabilidad.
* **Diseño Premium:** Estética de "Modo Oscuro" profesional optimizada con TailwindCSS.

### ⚙️ Módulo Administrativo (CRUD Web)
Interfaz exclusiva para administradores (`admin@stream.com`) que permite la gestión total del inventario sin tocar la base de datos directamente:
* **Crear:** Formulario dinámico para añadir películas con ID, título y descripción.
* **Leer:** Tabla de inventario que muestra todo el contenido cargado en MySQL.
* **Actualizar:** Sistema de edición mediante **ventanas modales** para modificar datos existentes en tiempo real.
* **Eliminar:** Opción de borrado permanente con confirmación de seguridad.

---

## 🛠️ Tecnologías Utilizadas

| Componente | Tecnología |
| :--- | :--- |
| **Backend** | Go (Golang) |
| **Base de Datos** | MySQL 8.0 |
| **Frontend** | HTML5, JavaScript (ES6+) |
| **Estilos** | TailwindCSS (vía CDN) |
| **Persistencia** | `database/sql` & `go-sql-driver/mysql` |

---

# 🛠️ Guía Técnica - StreamGo

Este documento detalla los requisitos, la configuración del entorno y los pasos necesarios para ejecutar el sistema de streaming de forma local.

---

## 📋 Requisitos Técnicos

Para ejecutar este proyecto, necesitas tener instalados los siguientes componentes:

1.  **Go (Golang):** Versión 1.18 o superior.
2.  **MySQL Server:** Versión 8.0 o superior.
3.  **Git:** Para la gestión del repositorio (opcional).
4.  **Navegador Web:** Chrome, Firefox o Edge.

---

## 🔧 1. Configuración de la Base de Datos

El sistema utiliza una base de datos MySQL. Sigue estos pasos para prepararla:

1. Crea una base de datos llamada `BDD_Streaming`.
2. Asegúrate de tener las tablas (`usuarios`, `perfiles`, `contenidos`, `planes_suscripcion`) creadas según el esquema del proyecto.
3. **Importante:** Ajusta las credenciales de conexión en el archivo `cmd/main.go` dentro de la función `main()`:

```go
s, err := storage.NewMySQLStorage(storage.DBConfig{
    User:     "root",             // Tu usuario de MySQL
    Password: "TU_PASSWORD_AQUÍ", // Tu contraseña de MySQL
    Host:     "localhost",
    Port:     "3306",
    DBName:   "BDD_Streaming",
})
```
## Ejecución del Proyecto

Abre una terminal en la raíz del proyecto y ejecuta los siguientes comandos para inicializar y arrancar el servidor:

```go
# Inicializar el módulo si no existe
go mod init streaming-system

# Descargar drivers de MySQL y dependencias
go mod tidy

# Ejecutar la aplicación
go run ./cmd/main.go
```
El sistema se iniciará en: http://localhost:8080.

## 🔒Acceso Administrativo

Para gestionar el contenido (CRUD), inicia sesión con la cuenta de administrador:

Email: admin@stream.com

Password: (La configurada en tu base de datos)

Nota: El sistema detecta automáticamente este correo y redirige al panel /admin.

## 📂 Estructura del Proyecto
streaming-system/
├── cmd/
│   └── main.go            # Servidor central y ruteo
├── pkg/
│   ├── auth/              # Usuarios y perfiles
│   ├── billing/           # Planes y suscripciones
│   ├── content/           # Modelos de catálogo
│   └── storage/           # Conector MySQL
├── templates/             # Vistas HTML (Login, Admin, Dashboard)
└── go.mod                 # Dependencias






