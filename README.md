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