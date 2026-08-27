# 🎟️ TicketEngine — High-Concurrency Ticketing Backend in Go

Backend distribuido de alto rendimiento diseñado en Go (Golang) bajo los principios de Clean Architecture. El sistema resuelve problemas de concurrencia y sobreventa en la compra masiva de boletos mediante contadores atómicos y bloqueos distribuidos en Redis, autenticación segura con JWT, y persistencia relacional con PostgreSQL + GORM.

---

## 🏗️ Estructura del Proyecto (Clean Architecture)

El proyecto está organizado siguiendo el patrón de capas desacopladas (domain, usecase, interface, infrastructure):

ticket-engine/
├── domain/
│   └── model/                  # Entidades de Dominio (User, Event, Reservation)
│       ├── event.go
│       ├── reservation.go
│       └── user.go
├── infrastructure/
│   ├── cache/                  # Conexión y scripts de Redis
│   ├── datastore/              # Conexión a PostgreSQL + GORM
│   ├── middleware/             # Middleware de Autenticación JWT
│   │   └── auth_middleware.go
│   └── router/                 # Configuración de rutas HTTP
│       ├── auth.go
│       ├── event.go
│       ├── reservation.go
│       ├── router.go
│       └── user.go
├── interface/
│   ├── controller/             # Controladores HTTP (Auth, User, Event, Reservation)
│   └── repository/             # Implementación concreta de Repositorios
├── usecase/
│   ├── interactor/             # Casos de Uso / Lógica de Negocio
│   └── repository/             # Interfaces de Repositorios (Ports)
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md

---

## 🗄️ Modelo de Datos Relacional

- User: Representa a los clientes de la plataforma. Mantiene el hash de contraseña (bcrypt) y la relación 1:N con las reservaciones (User posee varias Reservations).
- Event: Define el evento, precio por entrada, capacidad total y stock disponible en tiempo real.
- Reservation: Registra los boletos apartados por un usuario con estados (PENDING, CONFIRMED, EXPIRED, CANCELLED).

---

## ⚡ Manejo de Concurrencia, Autenticación y Prevención de Sobreventa

1. Autenticación JWT: El usuario se registra (/auth/register) o inicia sesión (/auth/login) para obtener un Bearer Token. El user_id se extrae de forma segura desde el token en el Middleware.
2. Atomic Decr (Redis): Cada solicitud descuenta primero el contador de boletos disponible en Redis. Si el conteo es < 0, la petición se rebota en menos de 5ms sin golpear la base de datos relacional.
3. Distributed Lock (10 min TTL): Se adquiere un bloqueo temporal por usuario/evento durante el proceso de pago. Si expira el tiempo sin confirmar el pago, el stock se reincorpora automáticamente.
4. Persistencia Transaccional (PostgreSQL + GORM): Cuando el pago se confirma, la reservación pasa a estado CONFIRMED dentro de una transacción en PostgreSQL.

---

## 📌 Endpoints API

### 🔑 Autenticación (Público)
- POST /auth/register - Registro de usuario (hash de contraseña con bcrypt)
- POST /auth/login - Login y obtención de Token JWT

### 🎟️ Eventos (Público / Admin)
- GET /events - Listar todos los eventos disponibles
- GET /events/:id - Consultar detalle y disponibilidad de boletos en tiempo real
- POST /events - Crear evento e inicializar stock atómico en Redis (Admin)

### 🔒 Reservaciones y Mis Entradas (Requiere Autenticación JWT)
- POST /reservations - Crear reservación atómica (Inicia TTL de 10 min en Redis)
- POST /reservations/:id/confirm - Confirmar pago y emitir boleto definitivo
- GET /users/me/reservations - Obtener el historial completo de reservaciones del usuario
- GET /users/me/tickets - Consultar únicamente los boletos confirmados/comprados