# Task Management API

A professional backend REST API built using **Go (Golang)**, **Gin**, **GORM**, **SQL Server**, and **Redis** following clean backend architecture principles.

This project includes JWT Authentication, Role-Based Authorization, CRUD Operations, Redis Caching, Docker, Swagger Documentation, Pagination, Search & Filtering.

---

# 🚀 Features

* JWT Authentication
* Role-Based Authorization
* CRUD Operations
* Repository Pattern
* Service Layer
* Middleware
* Redis Caching
* Pagination
* Search & Filtering
* Swagger Documentation
* Docker & Docker Compose
* SQL Server Integration
* Clean Architecture

---

# 🛠 Tech Stack

## Backend

* Go (Golang)
* Gin Framework
* GORM ORM

## Database

* SQL Server

## Cache

* Redis

## Authentication

* JWT Authentication

## Documentation

* Swagger

## DevOps

* Docker
* Docker Compose

---

# 📂 Project Structure

```text
task-management-api
│
├── cache
├── cmd
├── controllers
├── database
├── docs
├── dto
├── middleware
├── models
├── repositories
├── routes
├── services
├── utils
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── .env
```

---

# 🔐 Authentication

This API uses JWT Authentication.

Add token in request header:

```text
Authorization: Bearer YOUR_TOKEN
```

---

# 📘 Swagger Documentation

```text
http://localhost:8080/swagger/index.html
```

---

# 📌 API Endpoints

## Auth APIs

| Method | Endpoint           |
| ------ | ------------------ |
| POST   | /api/auth/register |
| POST   | /api/auth/login    |

---

## User APIs

| Method | Endpoint        |
| ------ | --------------- |
| GET    | /api/users      |
| GET    | /api/users/{id} |
| POST   | /api/users      |
| PUT    | /api/users/{id} |
| DELETE | /api/users/{id} |

---

## Task APIs

| Method | Endpoint        |
| ------ | --------------- |
| GET    | /api/tasks      |
| GET    | /api/tasks/{id} |
| POST   | /api/tasks      |
| PUT    | /api/tasks/{id} |
| DELETE | /api/tasks/{id} |

---

## Project APIs

| Method | Endpoint           |
| ------ | ------------------ |
| GET    | /api/projects      |
| GET    | /api/projects/{id} |
| POST   | /api/projects      |
| PUT    | /api/projects/{id} |
| DELETE | /api/projects/{id} |

---

# 📄 Pagination Example

```http
GET /api/tasks?page=1&limit=5
```

---

# 🔍 Search Example

```http
GET /api/tasks?search=API
```

---

# 🎯 Filter Example

```http
GET /api/tasks?status=Pending
```

---

# ⚡ Redis Caching

Task APIs use Redis caching for improving API performance and reducing database load.

---

# ▶️ Run Locally

## Clone Repository

```bash
git clone https://github.com/Dnyanesh0902/task-management-api.git
```

## Install Dependencies

```bash
go mod download
```

## Run Project

```bash
go run cmd/main.go
```

---

# 🐳 Run With Docker

```bash
docker compose up --build
```

---

# 🔧 Environment Variables

```env
DB_HOST=localhost
DB_PORT=1433
DB_USER=sa
DB_PASSWORD=Kokate@123
DB_NAME=taskmanagementapi

JWT_SECRET=your_secret_key

REDIS_ADDR=localhost:6379
```

---
# Task Management API Interview Questions & Answers

## 1. Explain your project architecture.

I used layered architecture in this project.

* Controllers handle HTTP requests and responses.
* Services contain business logic.
* Repositories handle database operations.
* Models define database schema.
* DTOs are used for request validation.
* Middleware handles JWT authentication and authorization.

This architecture improves maintainability, scalability, and code separation.

---

# 2. Why did you use Gin framework?

Gin is a lightweight and high-performance Go web framework.

Advantages:

* Fast routing
* Middleware support
* JSON handling
* Easy REST API development
* Better performance compared to many frameworks

---

# 3. Why did you use GORM?

GORM is an ORM library in Go.

Advantages:

* Simplifies database operations
* Auto migrations
* Relationship handling
* Query building
* Reduces SQL boilerplate code

---

# 4. What is JWT Authentication?

JWT (JSON Web Token) is used for secure authentication.

After login:

* Server generates token
* Client stores token
* Client sends token in Authorization header
* Middleware validates token

This enables stateless authentication.

---

# 5. Difference between Authentication and Authorization?

Authentication:

* Verifies who the user is

Authorization:

* Verifies what user can access

Example:

* Login is authentication
* Admin-only APIs are authorization

---

# 6. Why did you use Redis?

Redis is used for caching.

Benefits:

* Faster API response
* Reduced database load
* Better performance

I used Redis caching in GetTasks API.

---

# 7. What is Middleware in Gin?

Middleware executes before or after request processing.

In this project:

* JWT Middleware validates token
* Logger Middleware logs request details
* Role Middleware checks user roles

---

# 8. Why did you use Repository Pattern?

Repository pattern separates database logic from business logic.

Benefits:

* Clean code
* Easy testing
* Better maintainability
* Reusable database logic

---

# 9. Why did you use DTOs?

DTOs (Data Transfer Objects) are used for:

* Request validation
* Preventing direct model exposure
* Cleaner API contracts

---

# 10. What is Pagination?

Pagination divides large data into smaller pages.

Example:

```http
GET /api/tasks?page=1&limit=5
```

Benefits:

* Faster response
* Reduced memory usage
* Better frontend performance

---

# 11. How did you implement Search & Filter?

I used query parameters.

Examples:

```http
GET /api/tasks?search=API
```

```http
GET /api/tasks?status=Pending
```

This improves data filtering and user experience.

---

# 12. What is Docker?

Docker is a containerization platform.

Benefits:

* Same environment everywhere
* Easy deployment
* Simplified setup
* Better scalability

---

# 13. What is Docker Compose?

Docker Compose manages multiple containers together.

In this project:

* Go API container
* SQL Server container
* Redis container

All services run using single command:

```bash
docker compose up
```

---

# 14. Why Swagger is used?

Swagger provides API documentation.

Benefits:

* API testing
* Better developer experience
* Clear endpoint documentation
* Easy frontend integration

---

# 15. Explain Request Flow in Your Project.

Request Flow:

Client Request
↓
Routes
↓
Middleware
↓
Controller
↓
Service
↓
Repository
↓
Database

Response follows reverse flow.

---

# 16. Why Redis Cache Clear After Create/Update/Delete?

Because cached data becomes outdated after modifications.

So after:

* Create
* Update
* Delete

I clear cache to maintain fresh data consistency.

---

# 17. What is Preload in GORM?

Preload loads related data automatically.

Example:

```go
db.Preload("Project").Preload("AssignedUser")
```

It fetches related Project and User data with Task.

---

# 18. Difference Between PUT and PATCH?

PUT:

* Updates complete resource

PATCH:

* Updates partial resource

---

# 19. How did you secure passwords?

I used bcrypt hashing.

Passwords are never stored as plain text.

---

# 20. What improvements can be added in future?

Future improvements:

* CI/CD Pipeline
* Kubernetes
* Unit Testing
* Refresh Tokens
* Microservices
* gRPC
* Rate Limiting
* Email Notifications

---
# 👨‍💻 Author

## Dnyaneshwar Kokate

GitHub:
https://github.com/Dnyanesh0902

LinkedIn:
https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/

Portfolio:
https://dnyanesh.miracledevelopers.in