# Task Manager API

A lightweight RESTful API for task management built with Go. Designed with clean architecture principles, separation of concerns, and PostgreSQL database.

## Features

- ✅ Full CRUD operations for tasks (Create, Read, Update, Delete)
- ✅ PostgreSQL database with Docker
- ✅ pgAdmin web interface for database management
- ✅ Clean project architecture (models, handlers, connection)
- ✅ Proper HTTP status codes and error handling
- ✅ JSON request/response handling
- ✅ Health check endpoint
- ✅ Docker Compose for easy setup

## Quick Start

### Prerequisites

- Go 1.22+
- Docker & Docker Compose

### Installation

```bash
git clone https://github.com/yourusername/task-manager-api.git
cd task-manager-api
```

### Setup Database & Services

```bash
# Start PostgreSQL and pgAdmin with Docker
docker-compose up -d

# Initialize Go dependencies
go mod download
```

### Running the Application

```bash
go run main.go
```

Server akan start di `http://localhost:8080`

### Access Database Management

**pgAdmin Web Interface:**

- URL: http://localhost:5050
- Email: `admin@taskmanager.com`
- Password: `admin123`

Atau gunakan terminal untuk query:

```bash
docker exec -it taskmanager_postgres psql -U postgres -d taskmanager
```

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
