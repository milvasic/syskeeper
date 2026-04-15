# syskeeper

A Linux server monitoring and maintenance tool.

**Keeper** (server) — centralized .NET web app with PostgreSQL, providing agent registration, data storage, and a Razor Pages dashboard.

**Keepee** (agent) — lightweight Go binary installed on monitored machines, collecting system data (OS info, pending updates, Docker state) and pushing it to the keeper.

## Tech Stack

| Component     | Technology                         |
| ------------- | ---------------------------------- |
| Server        | ASP.NET Core + PostgreSQL + Dapper |
| UI            | Razor Pages                        |
| Agent         | Go                                 |
| Auth          | API key per agent                  |
| Communication | REST over HTTPS                    |

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

## License

MIT
