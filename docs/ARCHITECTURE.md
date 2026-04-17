# Architecture — syskeeper

This document describes the architecture of syskeeper: the keeper/keepee communication model, authentication, data flow, and the planned evolution toward OAuth2 and WebSocket-based instruction delivery.

---

## Table of Contents

- [Component Overview](#component-overview)
- [Keeper/Keepee Communication Model](#keeperkeepe-communication-model)
  - [Agent Registration](#agent-registration)
  - [Data Push (Ping)](#data-push-ping)
  - [Refresh Instruction](#refresh-instruction)
- [Data Flow Diagram](#data-flow-diagram)
- [Authentication Model](#authentication-model)
- [Planned: OAuth2 Flow](#planned-oauth2-flow)
- [Planned: WebSocket-Based Instruction Delivery](#planned-websocket-based-instruction-delivery)
- [Database Schema (Overview)](#database-schema-overview)
- [API Surface](#api-surface)

---

## Component Overview

| Component | Language / Framework | Role |
|-----------|----------------------|------|
| `keeper`  | ASP.NET Core 8, Razor Pages, Dapper, PostgreSQL | Central server: stores agent data, exposes REST API, serves dashboard |
| `keepee`  | Go 1.22+ | Lightweight agent: collects system info, registers with keeper, pushes data |

---

## Keeper/Keepee Communication Model

All communication is **HTTP/REST over HTTPS**. The `keepee` agent is the initiating party; the `keeper` server is passive (it does not initiate connections to agents).

### Agent Registration

```
keepee                        keeper
  │                              │
  │  POST /api/agents/register   │
  │  { hostname, os_info, ... }  │
  │ ────────────────────────────►│
  │                              │  Generate API key
  │  200 OK { api_key }          │  Store agent record
  │ ◄────────────────────────────│
  │                              │
```

The API key is stored by `keepee` locally and used for all subsequent requests.

### Data Push (Ping)

```
keepee                        keeper
  │                              │
  │  POST /api/agents/ping       │
  │  Authorization: Bearer <key> │
  │  { packages, docker, os }    │
  │ ────────────────────────────►│
  │                              │  Upsert snapshot in DB
  │  200 OK { instructions }     │  Return pending instructions
  │ ◄────────────────────────────│
  │                              │
```

`keepee` pings on a configurable interval (default: 5 minutes). Each ping carries the latest collected data.

### Refresh Instruction

The `keeper` can return instructions inside the ping response (e.g., `refresh_now`, `update_packages`). The `keepee` processes these instructions before sleeping until the next ping interval.

---

## Data Flow Diagram

```
┌────────────────────────────────────────────────────────────────┐
│  Monitored Hosts                                               │
│                                                                │
│  ┌──────────┐   ┌──────────┐   ┌──────────┐                  │
│  │ keepee-1 │   │ keepee-2 │   │ keepee-N │                  │
│  └────┬─────┘   └────┬─────┘   └────┬─────┘                  │
│       │              │              │                          │
└───────┼──────────────┼──────────────┼──────────────────────────┘
        │  HTTPS POST /api/agents/ping│
        └──────────────┴──────────────┘
                       │
                       ▼
        ┌──────────────────────────┐
        │         keeper           │
        │  (ASP.NET Core)          │
        │                          │
        │  ┌────────────────────┐  │
        │  │  Agent Controller  │  │
        │  └────────┬───────────┘  │
        │           │ Dapper       │
        │  ┌────────▼───────────┐  │
        │  │    PostgreSQL      │  │
        │  │  agents, snapshots │  │
        │  └────────────────────┘  │
        │                          │
        │  ┌────────────────────┐  │
        │  │  Razor Pages UI    │◄─┼── Browser (operator)
        │  └────────────────────┘  │
        └──────────────────────────┘
```

---

## Authentication Model

**Current:** Per-agent API keys.

- On registration, the `keeper` generates a cryptographically random API key and returns it once to the `keepee`.
- Subsequent requests carry the key in the `Authorization: Bearer <key>` header.
- Keys are stored hashed in PostgreSQL.
- There is no user authentication for the dashboard in the initial version; it is assumed to run on a private network.

---

## Planned: OAuth2 Flow

In a future iteration, the dashboard will support multi-user access via OAuth2/OIDC:

```
Browser                   keeper                    Identity Provider
   │                         │                              │
   │  GET /dashboard          │                              │
   │ ─────────────────────── ►│                              │
   │                          │  Redirect to /authorize      │
   │ ◄──────────────────────  │                              │
   │  GET /authorize          │                              │
   │ ──────────────────────────────────────────────────────► │
   │  code                    │                              │
   │ ◄──────────────────────────────────────────────────────│
   │  POST /token (code)      │                              │
   │ ──────────────────────────────────────────────────────► │
   │  { access_token }        │                              │
   │ ◄──────────────────────────────────────────────────────│
   │  GET /dashboard          │                              │
   │  Authorization: Bearer   │                              │
   │ ─────────────────────────►│                             │
   │  Dashboard HTML           │                              │
   │ ◄─────────────────────────│                             │
```

---

## Planned: WebSocket-Based Instruction Delivery

The current pull-based ping model has a latency equal to the ping interval. A planned enhancement will allow the `keeper` to push instructions to connected `keepee` agents in near-real time using WebSockets:

```
keepee                        keeper
  │                              │
  │  WS /api/agents/ws           │
  │  (persistent connection)     │
  │ ────────────────────────────►│
  │                              │
  │  ◄── { type: "ping_now" } ───│  (operator triggers from dashboard)
  │                              │
  │  POST /api/agents/ping       │
  │ ────────────────────────────►│
  │  ◄── 200 OK                  │
```

This model keeps the REST ping as the data transport while using WebSocket only for low-latency control signals.

---

## Database Schema (Overview)

```sql
-- Registered agents
agents (
  id          UUID PRIMARY KEY,
  hostname    TEXT NOT NULL,
  api_key_hash TEXT NOT NULL,
  registered_at TIMESTAMPTZ,
  last_seen_at  TIMESTAMPTZ
)

-- Latest snapshot per agent (upserted on each ping)
agent_snapshots (
  agent_id    UUID REFERENCES agents(id),
  os_info     JSONB,
  packages    JSONB,
  docker      JSONB,
  collected_at TIMESTAMPTZ,
  PRIMARY KEY (agent_id)
)
```

---

## API Surface

The full API specification is maintained in [`openapi/`](../openapi/). Key endpoints:

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/agents/register` | Register a new keepee agent |
| `POST` | `/api/agents/ping` | Push system snapshot, receive instructions |
| `GET`  | `/api/agents` | List all registered agents (dashboard use) |
| `GET`  | `/api/agents/{id}` | Get agent details and latest snapshot |
