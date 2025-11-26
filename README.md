# 🚀 Go E-Commerce Backend API

A robust and production-ready RESTful API for an E-Commerce platform, built with **Golang (Fiber)** using **Clean Architecture** principles. This project features secure authentication, transaction handling with data snapshots, and is fully containerized with **Docker**.

## 🛠️ Tech Stack

- **Language:** Golang 1.22+
- **Framework:** [GoFiber](https://gofiber.io/) (Fast HTTP Framework)
- **Database:** MySQL 8.0
- **ORM:** [GORM](https://gorm.io/)
- **Configuration:** [Viper](https://github.com/spf13/viper)
- **Logging:** [Logrus](https://github.com/sirupsen/logrus)
- **Migration:** [Golang-Migrate](https://github.com/golang-migrate/migrate)
- **Containerization:** Docker & Docker Compose

## ✨ Key Features

- **Authentication & Security**
  - Register & Login with JWT (JSON Web Token).
  - Secure Password Hashing (Bcrypt).
  - Role-Based Access Control (Admin & User Middleware).
- **User & Store Management**
  - Auto-create Store upon User Registration.
  - User Profile & Address Management.
- **Inventory Management**
  - Product CRUD with Image Upload.
  - Category Management (Admin only).
- **Transactions**
  - Checkout Logic with Stock Validation.
  - **Product Logging:** Snapshots product data (Price/Name) at the time of transaction into `log_products` table (Historical integrity).
  - Database Transactions.
- **Other**
  - Pagination & Filtering.
  - Standardized JSON Response.

## 📂 Project Structure (Clean Architecture)
```
.
├── cmd/                    # Application entrypoint
├── db/
│   └── migrations/         # Database migration files
├── internal/
│   ├── config/             # Viper config, DB, Logger, Fiber setup
│   ├── delivery/           # Handlers, Middleware
│   ├── entity/             # GORM models
│   ├── model/              # DTOs (Request/Response)
│   ├── repository/         # Database access layer
│   └── usecase/            # Business logic layer
├── public/                 # Uploaded asset files
├── docker-compose.yml
├── Dockerfile
└── config.example.json
```

---

## 🚀 How to Run (Docker Way - Recommended)

You don't need to install Go or MySQL on your local machine. Just use Docker!

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running.

### Steps
1. **Clone the repository**
   ```bash
   git clone https://github.com/arttVinci/Golang-API-Ecommerce.git
   cd Golang-API-Ecommerce
   ```

### **2. Jalankan dengan Docker Compose**
```bash
docker-compose up --build
```

API Base URL → `http://localhost:3000/api/ecommerce`

# 🧪 API Endpoints

Seluruh endpoint berada di prefix:

```
/api/ecommerce
```

---

# 🟦 Public Routes (No Auth Required)

### **Authentication**
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/register` | Register user baru + auto-create store |
| POST | `/login` | Login & generate JWT |

### **Products (Public Listing)**
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Search & List products (pagination, keyword, category) |

---

# 🟩 Protected Routes (Require JWT)

Prefix: `/api/ecommerce/*` dengan JWT middleware

### **User**
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users/current` | Get current user profile |
| PUT | `/users/current` | Update current user profile |

---

### **Store**
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/store` | Get my store |
| PUT | `/store` | Update store info |

---

### **Address**
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/addresses` | Create address |
| GET | `/addresses` | List addresses |
| DELETE | `/addresses/:id` | Delete address |

---

### **Products (Create Only, Auth Required)**
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/products` | Create new product (with image upload) |

---

### **Categories (Public list, rest admin only)**
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/categories` | List categories (public) |

---

### **Transactions**
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/transactions` | Checkout / Create transaction |
| GET | `/transactions` | Get transaction history |

---

# 🟥 Admin Routes (JWT + Admin Middleware)

Prefix: `/api/ecommerce/admin/*`

### **Admin Categories**
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/admin/categories` | Create category |
| PUT | `/admin/categories/:id` | Update category |
| DELETE | `/admin/categories/:id` | Delete category |

---

### Done! 🎉

## ⭐ Support

Kalau project ini membantu, jangan lupa kasih **star** di GitHub ⭐
