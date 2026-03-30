# Annora World Service

## Idea (Brief)

World Service is the backend service responsible for managing worlds in Annora.

At a high level, a world is the container where community-driven lore evolves over time. The service is intended to handle:

- world creation and discovery
- world membership and roles
- world-level metadata used by other domain flows (lore, voting, events, story phases)

Current implementation is focused on foundational backend layers (config, app wiring, HTTP scaffolding, models, repository, and migrations).

## � Hot Reload Development (Nodemon-like)

Air watches for changes to your Go files and automatically rebuilds and restarts the server. This is the Go equivalent of nodemon for Node.js.

### Installation

```bash
go install github.com/air-verse/air@latest
```

### Running with Auto-Reload

From the project root:

```bash
air
```

### How It Works

Air reads `.air.toml` which tells it to:
1. **Watch** for changes to `.go`, `.env`, and template files
2. **Ignore** the `tmp/` folder (where rebuilt binaries go), `vendor/`, and `_test.go` files
3. **Rebuild** by running: `go build -o ./tmp/main.exe ./cmd/world-service`
4. **Restart** the binary after rebuild
5. **Delay** 300ms before rebuild (waits for you to stop typing)
6. **Gracefully shut down** the old process before starting the new one (`send_interrupt`, `kill_delay`)
7. **Clear the terminal** on each rebuild so you see fresh output

The result: save a `.go` file → Air detects it → rebuilds in ~1 second → restarts the server automatically. No manual restart needed.

### Windows Path Note

If `air` command is not found, run:

```powershell
& "$env:USERPROFILE\go\bin\air.exe"
```

## �📁 Project Structure

```text
annora-world/
├── cmd/
│   └── annora-world/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   └── app.go
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── domain/
│   │   ├── model/
│   │   │   ├── world.go               ← World struct + methods
│   │   │   ├── membership.go          ← Membership struct, Role constants
│   │   │   └── timeline.go            ← TimelineEntry struct
│   │   ├── validation/
│   │   │   ├── world.go               ← ValidateWorld(), name rules etc.
│   │   │   └── membership.go          ← ValidateRole(), cap rules etc.
│   │   └── service/
│   │       ├── world.go               ← WorldService interface + struct
│   │       ├── membership.go          ← MembershipService interface + struct
│   │       └── timeline.go            ← TimelineService interface + struct
│   │
│   ├── api/
│   │   └── http/
│   │       ├── router.go
│   │       ├── middleware/
│   │       │   ├── auth.go
│   │       │   └── logger.go
│   │       ├── world/
│   │       │   ├── handler.go         ← create, get, update, delete world
│   │       │   ├── dto.go
│   │       │   └── routes.go
│   │       ├── membership/
│   │       │   ├── handler.go         ← join, leave, list members, update role
│   │       │   ├── dto.go
│   │       │   └── routes.go
│   │       └── timeline/
│   │           ├── handler.go         ← get timeline, get entry, list summaries
│   │           ├── dto.go
│   │           └── routes.go
│   │
│   ├── repository/
│   │   ├── world.go                   ← WorldRepository interface
│   │   ├── membership.go              ← MembershipRepository interface
│   │   └── timeline.go                ← TimelineRepository interface
│   │
│   └── infrastructure/
│       ├── postgres/
│       │   ├── client.go
│       │   ├── world.go               ← implements repository.WorldRepository
│       │   ├── membership.go          ← implements repository.MembershipRepository
│       │   └── timeline.go            ← implements repository.TimelineRepository
│       └── messaging/
│           └── publisher.go
│
├── migrations/
├── pkg/
│   └── response/
├── Makefile
├── Dockerfile
└── go.mod
```