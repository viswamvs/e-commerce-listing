# **E-Commerce Listing API 🚀**

This is a **Go-based e-commerce listing service** that provides CRUD operations for product management.  
It uses **Gin** for HTTP routing, **GORM** for database interactions, and **Redis** for caching.

---

## **📁 Project Structure**
```
e-commerce-listing/
│── daos/
│   ├── product/
│       ├── product.go         # DAO functions for DB operations
│
│── database/
│   ├── models/
│       ├── product.go         # Product model definition
│
│── dtos/
│   ├── product.go             # DTOs for API responses
│
│── handlers/
│   ├── product.go             # Handlers for CRUD operations
│
│── middleware/
│   ├── middleware.go          # Authentication middleware
│
│── routes/
│   ├── router.go              # Router setup
│   ├── product.go             # Product routes
│
│── services/
│   ├── product/
│       ├── product.go         # Business logic for product service
│
│── utils/
│   ├── cache/
│       ├── cache.go           # Redis cache initialization
│   ├── context/
│       ├── context.go         # Custom request context with Redis & DB
│   ├── db/
│       ├── db.go              # Database initialization & connection
│
│── vendor/                     # External dependencies (Go Modules)
│── README.md                   # Project Documentation
│── go.mod                      # Go module dependencies
│── go.sum                      # Dependency checksums
│── main.go                     # Entry point
```

---

## **🔧 Tech Stack**
- **Go (Golang)** 🏎️ - Backend language
- **Gin** 🌶️ - Web framework for routing & middleware
- **PostgreSQL** 🐘 - Primary database
- **GORM** 🏗️ - ORM for database interactions
- **Redis** ⚡ - In-memory caching for faster API responses

---

## **🚀 Setup & Installation**

### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/yourusername/e-commerce-listing.git
cd e-commerce-listing
```

### **2️⃣ Install Dependencies**
```sh
go mod tidy
```

### **5️⃣ Initialize Database**
```sh
go run main.go
```

---

## **🛠️ API Endpoints**

### **📌 Product Routes**
| Method | Endpoint               | Description               |
|--------|------------------------|---------------------------|
| GET    | `/v1/products`         | Get all products          |
| GET    | `/v1/products/:id`     | Get product by ID         |
| POST   | `/v1/products`         | Create a new product      |
| PUT    | `/v1/products/:id`     | Update a product          |
| DELETE | `/v1/products/:id`     | Delete a product          |

---

---

## **⚡ Redis Caching**
### **Caching Strategy**
- `GET /v1/products`:
  - If **cache hit**, return products from Redis.
  - If **cache miss**, fetch from DB and store in Redis for `10 minutes`.
- `POST / PUT`:
  - Update the **product key** in Redis.
  - Delete the `products` key to **invalidate the cache**.

### **Example Usage**
```go
c.Redis.Set(c.Request.Context(), "product:123", productJSON, 10*time.Minute)
c.Redis.Del(c.Request.Context(), "products")
```

---

## **🛠 Running the Server**
Start the server:
```sh
go run main.go
```
Your API will be available at:  
**📍 `http://localhost:8080`**

---

## **📝 TODO**
- ✅ Implement **CRUD APIs**
- ✅ Add **Redis caching**
- ⏳ Use **middleware for authentication**
- ⏳ Implement **search & filtering**
- ⏳ Add **unit tests & integration tests**

---

## **📜 License**
This project is licensed under **MIT License**.

---
