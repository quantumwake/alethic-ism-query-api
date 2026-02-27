# Alethic ISM Query API

Go/Gin service exposing state query, vault, and NLP embedding endpoints for the Alethic ISM platform. Each route group maps to a distinct domain for clean ingress-to-service routing.

## Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| POST | `/api/v1/state/query/:id` | Query state data with DSL filters |
| POST | `/api/v1/vault` | Create a vault |
| GET | `/api/v1/vault/:id` | Fetch a vault |
| DELETE | `/api/v1/vault/:id` | Delete a vault |
| POST | `/api/v1/nlp/embeddings` | Upsert an embedding document |
| POST | `/api/v1/nlp/embeddings/batch` | Batch upsert documents |
| GET | `/api/v1/nlp/embeddings/:id` | Get document by ID |
| GET | `/api/v1/nlp/embeddings/parent/:id` | Get documents by parent ID |
| DELETE | `/api/v1/nlp/embeddings/:id` | Delete document by ID |
| DELETE | `/api/v1/nlp/embeddings/parent/:id` | Delete documents by parent ID |
| POST | `/api/v1/nlp/embeddings/search` | Similarity search |
| POST | `/api/v1/nlp/embeddings/migrate` | Run table migration |

Swagger UI at `/swagger/index.html`, ReDoc at `/redoc/index.html`.

## Setup

Requires Go 1.24+ and PostgreSQL with pgvector.

| Env Var | Default | Description |
|---------|---------|-------------|
| `DSN` | — | PostgreSQL connection string |
| `PORT` | `8080` | HTTP listen port |

```bash
go mod download
go build -o main .
./main
```

## Deployment

K8s manifests in `k8s/`. Ingress routes `/api/v1/state/query`, `/api/v1/vault`, and `/api/v1/nlp/embeddings` to this service.

## License

Alethic ISM is under a DUAL licensing model, please refer to [LICENSE.md](LICENSE.md).

**AGPL v3**  
Intended for academic, research, and nonprofit institutional use. As long as all derivative works are also open-sourced under the same license, you are free to use, modify, and distribute the software.

**Commercial License**
Intended for commercial use, including production deployments and proprietary applications. This license allows for closed-source derivative works and commercial distribution. Please contact us for more information.