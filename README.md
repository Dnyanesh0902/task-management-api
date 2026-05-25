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

# 👨‍💻 Author

## Dnyaneshwar Kokate

GitHub:
https://github.com/Dnyanesh0902

LinkedIn:
https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/

Portfolio:
https://dnyanesh.miracledevelopers.in