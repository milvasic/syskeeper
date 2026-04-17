# syskeeper

> A Linux server monitoring and maintenance tool — a lightweight, self-hosted alternative to commercial fleet management.

**Keeper** (server) — centralized ASP.NET Core web application with PostgreSQL that provides agent registration, system data storage, and a Razor Pages dashboard.

**Keepee** (agent) — lightweight Go binary installed on each monitored Linux host; collects system information (OS details, pending package updates, Docker container state) and pushes it to the keeper.

---

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Monorepo Layout](#monorepo-layout)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Run the keeper](#run-the-keeper)
  - [Run a keepee agent](#run-a-keepee-agent)
- [Contributing](#contributing)
- [Security](#security)
- [License](#license)

---

## Overview

syskeeper lets you monitor a fleet of Linux servers from a single dashboard. Each server runs a `keepee` agent that periodically pushes:

- OS information (distribution, kernel version, hostname)
- Pending package updates
- Docker container states

All data is stored in the `keeper` server and displayed in a Razor Pages web UI.

---

## Architecture

```
┌─────────────┐       HTTPS/REST        ┌─────────────────────┐
│   keepee    │ ──── push data ────────► │      keeper         │
│  (Go agent) │ ◄─── refresh request ── │  (ASP.NET Core)     │
│             │                          │  ┌───────────────┐  │
│ collects:   │      register            │  │ Razor Pages   │  │
│ - packages  │ ────────────────────────►│  │ Dashboard     │  │
│ - OS info   │      ◄── API key ──────  │  └───────────────┘  │
│ - docker    │                          │  ┌───────────────┐  │
└─────────────┘                          │  │ PostgreSQL    │  │
                                         │  └───────────────┘  │
                                         └─────────────────────┘
```

See [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) for a detailed description of the keeper/keepee communication model, data flow, and planned OAuth2/WebSocket enhancements.

---

## Tech Stack

| Component     | Technology                         |
| ------------- | ---------------------------------- |
| Server        | ASP.NET Core + PostgreSQL + Dapper |
| UI            | Razor Pages                        |
| Agent         | Go                                 |
| Auth          | API key per agent                  |
| Communication | REST over HTTPS                    |

---

## Monorepo Layout

```
syskeeper/
├── keeper/          # ASP.NET Core server (Razor Pages + PostgreSQL + Dapper)
├── keepee/          # Go agent binary
├── docs/            # Architecture documentation
│   └── ADR/         # Architecture Decision Records
├── openapi/         # OpenAPI specification for the keeper API
├── .github/         # GitHub Actions, issue/PR templates, Copilot instructions
├── AI_GUIDELINES.md
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
├── LICENSE
├── README.md
└── SECURITY.md
```

---

## Getting Started

### Prerequisites

| Tool | Minimum version | Purpose |
|------|-----------------|---------|
| .NET SDK | 8.0 | Build and run `keeper` |
| Go | 1.22 | Build and run `keepee` |
| PostgreSQL | 15 | `keeper` database |
| Docker (optional) | 24 | Run PostgreSQL locally |

### Run the keeper

```bash
cd keeper
cp appsettings.json appsettings.Development.json
# Edit appsettings.Development.json — set your PostgreSQL connection string

dotnet restore
dotnet run
```

The dashboard will be available at `http://localhost:5000`.

### Run a keepee agent

```bash
cd keepee
go mod download
go run . --keeper-url http://localhost:5000 --api-key <your-key>
```

On first run, `keepee` registers itself with the `keeper` server and receives an API key for subsequent communication.

---

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for:

- Branching strategy and commit conventions (Conventional Commits)
- Pull request process and review expectations
- Coding guidelines for both .NET and Go components
- AI-assisted contribution policy ([AI_GUIDELINES.md](AI_GUIDELINES.md))

---

## Security

Please **do not** open public GitHub issues for security vulnerabilities. See [SECURITY.md](SECURITY.md) for the responsible disclosure process.

---

## License

MIT — see [LICENSE](LICENSE).
