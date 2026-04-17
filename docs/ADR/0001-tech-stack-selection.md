# ADR 0001 — Tech Stack Selection

**Date:** 2025-04-16  
**Status:** Accepted  
**Deciders:** milvasic  
**Ticket / Issue:** [#25 — Project governance & community files](https://github.com/milvasic/syskeeper/issues/25)

---

## Context

syskeeper is a self-hosted Linux server monitoring tool. The project requires two main components:

1. A **central server** that stores agent data, provides a web dashboard, and exposes a REST API.
2. A **lightweight agent binary** installed on each monitored Linux host that collects system information and pushes it to the server.

The key constraints for the tech stack are:

- The author's existing proficiency and preference.
- The agent must be **easy to distribute** — a single static binary is strongly preferred to avoid dependency management on monitored hosts.
- The server should support **relational data storage** for structured agent snapshots and **a web UI** without a separate front-end build pipeline.
- The project is personal / open-source, so the stack should have strong community support and free tooling.

---

## Decision

**Server (`keeper`):** ASP.NET Core (C#) with Razor Pages, PostgreSQL, and Dapper.  
**Agent (`keepee`):** Go.

---

## Considered Alternatives

| Alternative | Reason not chosen |
|-------------|-------------------|
| Python (FastAPI + SQLAlchemy) for the server | Less familiar to the author; slower cold starts for a long-running server process are not a concern, but the ecosystem for web UIs without a JS build step is weaker. |
| Node.js / Express for the server | Author preference for statically typed languages in server code. |
| Rust for the agent | Steeper learning curve; Go provides sufficient performance and a simpler cross-compilation story for the target use case. |
| SQLite for the database | PostgreSQL is preferred for its robustness, JSONB support for flexible snapshot storage, and production-readiness. |
| Entity Framework Core for ORM | Dapper is lighter weight, gives more control over SQL, and is better suited to the simple, query-oriented data access pattern of this project. |
| Full SPA (React/Vue) for the dashboard | A separate front-end build pipeline adds complexity without clear benefit for a personal monitoring tool. Razor Pages provides server-rendered HTML with minimal overhead. |

---

## Consequences

### Positive

- **Go agent** compiles to a single static binary, making installation on monitored hosts trivial (`scp` + `chmod +x`).
- **ASP.NET Core** has first-class async support, strong type safety, and excellent performance.
- **Razor Pages** delivers a functional web UI without requiring a separate JS framework or build toolchain.
- **Dapper** keeps database access explicit and easy to understand.
- **PostgreSQL** provides JSONB columns for flexible storage of system snapshots without requiring rigid migrations every time a new data field is added.

### Negative / Trade-offs

- The monorepo contains two separate language ecosystems (.NET and Go), which increases the complexity of CI/CD pipelines and local development setup.
- Contributors need familiarity with both C# and Go to work across the full stack.
- Razor Pages does not support reactive/real-time UI updates without additional work (e.g., SignalR or polling).

### Neutral

- Both ecosystems have mature container support, so Dockerisation is straightforward when needed.
- Cross-compilation of the Go agent binary for multiple Linux architectures (amd64, arm64) is supported out of the box.

---

## References

- [ASP.NET Core documentation](https://learn.microsoft.com/en-us/aspnet/core/)
- [Go documentation](https://go.dev/doc/)
- [Dapper on GitHub](https://github.com/DapperLib/Dapper)
- [PostgreSQL JSONB](https://www.postgresql.org/docs/current/datatype-json.html)
