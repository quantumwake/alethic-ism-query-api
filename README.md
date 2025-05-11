# Alethic Instruction-Based State Machine (State Query API)

A RESTful API service for querying state data from the Alethic Instruction-Based State Machine system. This API provides endpoints for querying state entries based on various criteria and managing vaults.

## Features

- **State Querying**: Query state data using a flexible DSL (Domain Specific Language)
- **Vault Management**: Create, retrieve, and delete vaults
- **Swagger Documentation**: Comprehensive API documentation via Swagger UI and ReDoc
- **PostgreSQL Integration**: Persistence layer using PostgreSQL with GORM

## API Endpoints

### State Query

- `POST /api/v1/state/{id}/query`: Query state data with filters

### Vault Management

- `POST /api/v1/vault`: Create a new vault
- `GET /api/v1/vault/{id}`: Retrieve vault data by ID
- `DELETE /api/v1/vault/{id}`: Delete a vault by ID

## Documentation

API documentation is available at:

- Swagger UI: `/swagger/index.html`
- ReDoc: `/redoc/index.html`

## Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL database

### Environment Variables

- `DSN`: Database connection string (default: `host=localhost port=5432 user=postgres password=postgres1 dbname=postgres sslmode=disable`)

### Installation

1. Clone the repository
2. Install dependencies with `go mod download`
3. Build the API documentation with `./buildswag.sh`
4. Build the application with `go build -o main .`
5. Run the application with `./main`

### Docker

Build and run using Docker:

```bash
docker build -t alethic-ism-query-api --build-arg GIT_USERNAME=your_username --build-arg GIT_TOKEN=your_token .
docker run -p 8081:8081 alethic-ism-query-api
```

## Deployment

Kubernetes deployment configuration is available in the `k8s` directory.

## License

Alethic ISM is under a DUAL licensing model, please refer to [LICENSE.md](LICENSE.md).

**AGPL v3**  
Intended for academic, research, and nonprofit institutional use. As long as all derivative works are also open-sourced under the same license, you are free to use, modify, and distribute the software.

**Commercial License**
Intended for commercial use, including production deployments and proprietary applications. This license allows for closed-source derivative works and commercial distribution. Please contact us for more information.