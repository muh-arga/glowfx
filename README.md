# GlowFX - Fullstack Product Management

built with:
* Backend: Golang (Gin + GORM)
* Frontend: React (Vite)
* Database: PostgreSQL

---

## Features

* View product list
* Search & filter
* Pagination
* Create product
* Update product
* Delete product
* Validation (including SKU unique)

---

## Project Structure

```
glowfx/
├── backend/
├── frontend/
```

---

# Getting Started

## 1. Clone Project

```bash
git clone https://github.com/muh-arga/glowfx.git
cd glowfx
```

---

# Backend Setup (Golang)

## Masuk ke folder backend

```bash
cd backend
```

---


# Environment Setup

Project ini menggunakan file `.env` untuk konfigurasi.

## 1. Copy dari example

```bash
cp .env.example .env
```

---

## 2. Isi konfigurasi

Edit ``` 
    DB_USER, DB_PASSWORD, DB_NAME,
``` pada file `.env`:

---

## Penjelasan

| Variable    | Description            |
| ----------- | ---------------------- |
| APP_ENV     | Environment (dev/prod) |
| APP_PORT    | Port backend           |
| DB_HOST     | Database host          |
| DB_PORT     | Database port          |
| DB_USER     | Database user          |
| DB_PASSWORD | Database password      |
| DB_NAME     | Database name          |

---

## Setup Database (PostgreSQL)

Buat database:

```sql
CREATE DATABASE glowfx;
```

---

## Install Dependency

```bash
go mod tidy
```

---

## Run Backend

```bash
go run main.go
```

Server akan jalan di:

```
http://localhost:8085
```

---

## Note

* Database akan otomatis migrate saat aplikasi start (`AutoMigrate`)

---

# Frontend Setup (React)

## Masuk ke folder frontend

```bash
cd ../frontend
```

---

# Environment Setup

Project ini menggunakan file `.env` untuk konfigurasi.

## 1. Copy dari example

```bash
cp .env.example .env
```

---

## 2. Isi konfigurasi

Edit ``` 
    VITE_API_URL
``` pada file `.env` sesuai dengann URL backend:

---

## Penjelasan

| Variable      | Description       |
| -----------   | ------------------|
| VITE_API_URL  | URL Backend       |

---

## Install Dependency

```bash
npm install
```

---

## Run Frontend

```bash
npm run dev
```

Akses di browser:

```
http://localhost:5173
```

---

# API Endpoint

Base URL:

```
http://localhost:8085/api
```

### Product

| Method | Endpoint      | Description |
| ------ | ------------- | ----------- |
| GET    | /products     | Get list    |
| GET    | /products/:id | Get detail  |
| POST   | /products     | Create      |
| PUT    | /products/:id | Update      |
| DELETE | /products/:id | Delete      |

---

# Example Request

### Create Product

```json
{
  "Name": "Lipstick",
  "SKU": "LS-001",
  "Status": "active"
}
```

---
