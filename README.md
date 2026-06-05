# Task Manager API

A lightweight RESTful API for task management built with Go. Designed with clean architecture principles, separation of concerns, and concurrent-safe data operations.

## Features

- ✅ Full CRUD operations for tasks
- ✅ Thread-safe in-memory storage
- ✅ Clean project architecture (models, handlers, main)
- ✅ Proper HTTP status codes and error handling
- ✅ JSON request/response handling
- ✅ Health check endpoint

## Quick Start

### Prerequisites
- Go 1.22+

### Installation

```bash
git clone https://github.com/yourusername/task-manager-api.git
cd task-manager-api
go mod download
```

### Running the Server

```bash
go run main.go
```

Server will start at `http://localhost:8080`

## API Examples

**Create a task:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Master REST API development","completed":false}'
```

**Get all tasks:**
```bash
curl http://localhost:8080/tasks
```

**Get specific task:**
```bash
curl http://localhost:8080/tasks/1
```

**Update a task:**
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"REST API mastery","completed":true}'
```

**Delete a task:**
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Architecture

### Separation of Concerns
- **models/**: Business logic and data storage
- **handlers/**: HTTP request/response handling
- **main.go**: Server configuration and routing

### Key Concepts
- RESTful API design patterns
- HTTP method-based routing
- JSON serialization/deserialization
- Thread-safe concurrent access (sync.RWMutex)
- Proper error handling and HTTP status codes
- Clean code organization

## Testing

### Using Postman
1. Import `API.json` into Postman
2. Execute requests against running server

### Using curl
See "API Examples" section above

## Technical Stack

- **Language**: Go 1.22
- **Architecture**: Clean Architecture with separation of concerns
- **Concurrency**: sync.RWMutex for thread-safe operations
- **Storage**: In-memory (map-based)

## Future Enhancements

- Database integration (PostgreSQL/SQLite)
- Request validation middleware
- User authentication & authorization
- Comprehensive unit tests
- API documentation (Swagger/OpenAPI)
- Graceful shutdown handling
- Structured logging

## License

MIT
