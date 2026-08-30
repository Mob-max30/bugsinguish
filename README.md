# Bugsinguish — Pavan's Backend (feature/pavan-api)

## Setup
```
cd backend
go mod tidy      # resolves exact dependency versions from go.mod
```

Create a `.env` (not committed) at the repo root with:
```
DATABASE_URL=postgres://...neon-connection-string...
GEMINI_API_KEY=your-key-here
```

Run locally:
```
docker-compose up --build
```
or just the backend directly:
```
cd backend
go run main.go
```

## File map
- `backend/main.go` — Chi router, health check, wires up ticket + SSE routes, connects to Neon on boot
- `backend/models/ticket.go` — Ticket struct (post this shape to the team channel)
- `backend/db/neon.go` — Postgres/pgvector connection + tickets table migration
- `backend/db/dedup.go` — pgvector cosine-distance duplicate lookup
- `backend/handlers/tickets.go` — CRUD endpoints, with embedding + dedup-check wired into POST /tickets
- `backend/embeddings/embed.go` — Gemini embedding wrapper ([]float32, 768-dim)
- `backend/sse/stream.go` — GET /api/stream + in-memory pub/sub (post the Event JSON shape to the team channel)
- `docker-compose.yml` — frontend + backend, Docker socket mounted for Ujwal's sandbox manager

## Still open
- Item 9 (dependent): endpoint to receive Ujwal's RCA/CrashReport output and publish
  the "diagnosed" SSE event — needs his JSON shape first.
