# Armis

Armis is a lightweight, fully distributed key-value database with put, get, and delete operations. It supports CLI, HTTP API (with Swagger), and upcoming gRPC. Built in Go with in-memory storage (persistent/distributed features in development).

## Features

- Key-value CRUD: put, get, delete.
- Interfaces: CLI, HTTP (Swagger docs), gRPC (soon).
- Distributed design (initial in-memory).
- Dockerized for easy deployment.

## Installation

### From Source
```
git clone https://github.com/DKeshavarz/armis.git
cd armis
go mod tidy
go build -o armis ./cmd/main.go
```

### Docker
```
docker pull ghcr.io/dkeshavarz/armis:latest
```

## Usage

### Running
- Go: `go run ./cmd/main.go` (starts HTTP server and CLI).
- Docker: `docker run -p 8080:8080 ghcr.io/dkeshavarz/armis:latest` (HTTP at localhost:8080). For CLI: `docker run -it ghcr.io/dkeshavarz/armis:latest`.

### Configuration
Uses `.env` 
```bash
# Server Configuration
PORT=8080

# Storage Configuration
STORAGE_AUTO_SAVE=true
STORAGE_SAVE_INTERVAL=3
STORAGE_FILE_PATH=storage.json
```

### HTTP API
Endpoints at `/client`. Docs: http://localhost:8080/swagger/index.html.

- PUT `/client/:key` (body: value).
- GET `/client/:key`.
- DELETE `/client/:key`.

### CLI
Interactive prompt. See [CLI README](./internal/commands/README.md) for details.

## Project Structure
- `cmd/`: App entry.
- `internal/commands/`: CLI handlers.
- `internal/server/`: HTTP setup.
- `internal/service/`: Logic.
- `internal/storage/`: Storage.
- `config/`: Env loading.

## Contributing
Fork, branch, commit, PR. Issues: [GitHub](https://github.com/DKeshavarz/armis/issues).

## License
MIT. See [LICENSE](LICENSE).

## Contact
- GitHub: [DKeshavarz/armis](https://github.com/DKeshavarz/armis)
