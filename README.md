# 🎬 Streaming Management System (Core Backend)

## 🔰 Estado del Proyecto

| Fase Actual | Estado | Descripción |
| :--- | :--- | :--- |
| **Planeación (Etapa 1)** | ✅ **COMPLETADA** | Se han definido el objetivo, módulos, estructura de directorios, entidades clave y dependencias de terceros. |
| Implementación (Etapa 2) | ⏸ Pendiente | Desarrollo del código en Go, implementando la lógica funcional y la conexión a la base de datos. |

---

## 📌 1. Introducción

Este repositorio contiene el sistema de gestión central (*backend/core*) para una plataforma de distribución de contenido multimedia bajo demanda (Video on Demand, VoD).

El proyecto se enmarca en un ejercicio de diseño de software, donde la planeación arquitectónica se basa en un **modelado de entidades de POO (Diagrama de Clases)**, mientras que la implementación futura se realizará bajo el paradigma de **Programación Funcional** utilizando el lenguaje **Go**. Este enfoque dual busca un diseño estructurado y una implementación orientada al rendimiento y la escalabilidad.

---

## ✨ 2. Módulos Principales (Lógica de Negocio)

El sistema se organiza en torno a cuatro módulos funcionales que componen el *core* de la plataforma:

1.  **I. Gestión de Usuarios y Autenticación:** Manejo de credenciales, perfiles y roles.
2.  **II. Gestión de Contenido (Catálogo):** CRUD para títulos (películas, series), búsqueda y clasificación.
3.  **III. Gestión de Suscripciones y Pagos:** Administración de planes de servicio y el estado de la membresía.
4.  **IV. Reproducción y Historial (Simulado):** Registro de la actividad de consumo (progreso de visualización).

---

## 🏗 3. Estructura del Directorio (Go Standard Layout)

La arquitectura sigue las convenciones estándar de proyectos en Go para una clara separación de responsabilidades:

## ⚙️ 4. Tecnologías y Dependencias

| Componente | Tecnología/Paquete | Razón de Uso |
| :--- | :--- | :--- |
| **Lenguaje** | **Go (Golang)** | Rendimiento, concurrencia y enfoque en la simplicidad para el desarrollo *backend*. |
| **Base de Datos** | [Sugiera una, ej: **PostgreSQL**] | Robustez, integridad de datos y soporte avanzado para estructuras relacionales. |
| **Framework Web** | `github.com/gin-gonic/gin` | Router $\text{HTTP}$ ligero y de alto rendimiento para construir la $\text{API}$ $\text{RESTful}$. |
| **Conector DB** | `gorm.io/gorm` o `github.com/lib/pq` | Gestión eficiente de la conexión y las consultas a la base de datos. |
| **Seguridad** | `golang.org/x/crypto/bcrypt` | Implementación estándar y segura para el *hashing* de contraseñas de usuario. |

---

## 📚 5. Diagrama de Entidades Clave (Modelado POO)

El modelado define las relaciones fundamentales que guiarán las estructuras de datos en Go:

* **Entidades Centrales:** `Usuario`, `Suscripción`, `Contenido`, `Historial`, `Transacción`.
* **Relaciones:**
    * **Usuario** tiene una **Suscripción** ($1:1$).
    * **Usuario** está relacionado con **Contenido** a través de **Historial** ($\text{N}:\text{M}$).

**(Nota: El diagrama visual de clases debe ser añadido al documento $\text{PDF}$ y se sugiere incluir una imagen aquí en el $\text{README}$ en etapas posteriores.)**

---

## 🚀 Guía de Instalación (Próximas Fases)

1.  **Clonar el repositorio:**
    ```bash
    git clone [https://github.com/](https://github.com/)[TuUsuario]/streaming-management-system.git
    cd streaming-management-system
    ```
2.  **Inicializar Módulos de Go:**
    ```bash
    go mod tidy
    ```
3.  **Ejecutar el Core:**
    ```bash
    go run cmd/main.go
    ```


