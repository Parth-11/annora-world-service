# Annora World Service

## Idea (Brief)

World Service is the backend service responsible for managing worlds in Annora.

At a high level, a world is the container where community-driven lore evolves over time. The service is intended to handle:

- world creation and discovery
- world membership and roles
- world-level metadata used by other domain flows (lore, voting, events, story phases)

Current implementation is focused on foundational backend layers (config, app wiring, HTTP scaffolding, models, repository, and migrations).

## 📁 Project Structure

```text
annora-world-service/
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── cmd/
│   └── world-service/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── config/
│   │   ├── cache.go
│   │   ├── config.go
│   │   ├── env.go
│   │   ├── db.go
│   │   ├── queue.go
│   │   └── server.go
│   ├── dtos/
│   ├── handlers/
│   │   └── https/
│   │       ├── router.go
│   │       └── world/
│   ├── infrastructure/
│   │   ├── cache/
│   │   ├── db/
│   │   └── queue/
│   ├── models/
│   ├── repository/
│   │   └── world/
│   │   └── memberships/
│   └── service/
│       └── world/
│       └── memberships/
└── migrations/
```