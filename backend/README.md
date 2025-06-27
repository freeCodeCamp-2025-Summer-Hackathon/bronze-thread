## Installation Setup

1. Create a `.env` file - use `.env.example` for reference.

2. Run the server:

```bash
go run cmd/server/main.go
```

The server will start at http://localhost:8080 or on PORT in your `.env` file.

## Folder structure

```
/backend/cmd/server
│
├── /controllers       # Handler functions
│   └── <controller_name>_controller.go
│
├── /routes            # Route definitions
│   └── <route_name>_routes.go
│
├── /services          # Business logic
│   └── <service_name>_service.go
│
└── main.go            # App entry point
```
