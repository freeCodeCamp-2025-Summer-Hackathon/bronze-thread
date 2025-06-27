## Installation Setup

1. Create a `.env` file - use `.env.example` for reference.

2. Run the server:

```bash
go run cmd/server/main.go
```

Or with auto-reload:

```bash
air     # make sure to install air and configure .air.toml file
```

The server will start at http://localhost:8080 or on PORT in your `.env` file.

## Prerequisites

- [Go](https://go.dev)
- [Air](https://github.com/air-verse/air) (for reload on file changes)
