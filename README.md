# Go Project Template

Template for Go backend application.

This is a boilerplate template for building scalable and maintainable Go applications 
using Domain-Driven Design (DDD) principles. It provides a well-structured project layout 
and best practices for separating concerns and managing dependencies.

# Project Structure

```text
├── api/                # API related documentation.
│   └── docs/           # Generated swagger documentation.
├── cmd/                # Entry points for the application (e.g., HTTP server, CLI, etc.)
├── domain/             # Core domain logic (Entities, Value Objects, Aggregates, Interfaces)
│   ├── greet/          # Greet domain
│   └── ...             # Another domain
├── infra/              # Frameworks, database, and external APIs
│   ├── storage/        # Database implementation (Postgres, Redis, etc.)
│   ├── http/           # HTTP server (Handlers, Routers)
│   └── ...             # Another domain
├── pkg/                # Shared libraries or utilities
├── script/             # Utility scripts
├── .gitignore          # gitignore file
├── Dockerfile          # Application Dockerfile
├── go.mod              # Go module definition
├── Makefile            # Makefile
└── README.md           # Project documentation
```
