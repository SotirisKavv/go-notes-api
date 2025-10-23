# 📝 Notes API (Go)

A production-ready REST API for note management with SQLite/in-memory storage, middleware architecture, and authentication. Demonstrates advanced Go web development patterns, database integration, and enterprise-level API design.

---

## 🚀 What is this?

A fully-featured REST API showcasing enterprise patterns including middleware chains, repository abstraction, database integration with SQLite, and authentication. Perfect for demonstrating production-ready Go web development skills.

---

## ✨ Features

- **Complete CRUD API:** Create, read, update, delete notes with RESTful endpoints
- **Dual Storage:** SQLite database or in-memory storage (configurable)
- **Middleware Architecture:** Authentication and logging middleware chains
- **Database Integration:** SQLite with proper connection management
- **Authentication:** Token-based auth system with middleware
- **Error Handling:** Comprehensive HTTP status codes and error responses

---

## 🦄 Go Skills Demonstrated

- **Advanced Web Development:** RESTful API design with `gorilla/mux`
- **Middleware Patterns:** Custom authentication and logging middleware
- **Database Programming:** SQLite integration with proper error handling
- **Repository Pattern:** Interface-based storage abstraction
- **HTTP Best Practices:** Status codes, headers, and JSON responses
- **Dependency Injection:** Clean architecture with testable components

---

## 🛠️ Usage

```sh
# Install dependencies
go mod tidy

# Run with SQLite (requires CGO)
$env:CGO_ENABLED=1  # Windows PowerShell
go run .

# Test API endpoints
curl -Method POST http://localhost:8080/note \
  -Headers @{ "Authorization" = "Bearer mytoken"; "Content-Type" = "application/json" } \
  -Body '{ "title": "Todo", "body": "Learn Go", "user_id": 1 }'
```

**API Endpoints:**
- `GET /notes` - List all notes (auth required)
- `GET /note/{id}` - Get specific note (auth required)  
- `POST /note` - Create new note (auth required)
- `PUT /note/{id}` - Update note (auth required)
- `DELETE /note/{id}` - Delete note (auth required)

---

## 🎯 Learning Objectives

This project demonstrates:
- **Enterprise API Design:** Production-ready REST API architecture
- **Database Integration:** SQLite with proper connection handling
- **Middleware Development:** Custom auth and logging middleware chains
- **Clean Architecture:** Repository pattern and dependency injection

---

**Author:** IAmSotiris
