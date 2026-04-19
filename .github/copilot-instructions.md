# GitHub Copilot Instructions for syskeeper

This file provides context for GitHub Copilot to generate more relevant suggestions for the syskeeper project.

---

## Project Overview

**syskeeper** is a self-hosted Linux server monitoring tool consisting of two components:

- **`keeper`** — ASP.NET Core 10 server with Razor Pages dashboard, PostgreSQL (via Dapper), REST API.
- **`keepee`** — Go 1.26+ agent binary that runs on monitored Linux hosts and pushes system data to the keeper.

---

## Repository Structure

```
syskeeper/
├── keeper/          # C# / ASP.NET Core server
├── keepee/          # Go agent
├── docs/ARCHITECTURE.md
├── docs/ADR/        # Architecture Decision Records
├── openapi/         # OpenAPI spec
└── .github/         # Templates and this file
```

---

## Coding Conventions

### keeper (C# / ASP.NET Core)

- **Namespace style:** `Keeper.<Layer>` (e.g., `Keeper.Api`, `Keeper.Data`, `Keeper.Pages`)
- **Async everywhere:** All I/O methods must be `async Task<T>`. Never use `.Result` or `.Wait()`.
- **Data access:** Use Dapper; no Entity Framework. Repository classes live in `Keeper.Data`.
- **Validation:** Use `DataAnnotations` or FluentValidation on input models; never trust agent-supplied data.
- **Error handling:** Use `ProblemDetails` (RFC 7807) for API error responses.
- **Logging:** Use `ILogger<T>` with structured log messages. Avoid string interpolation in log calls; use message templates.
- **Naming:**
  - Controllers: `<Resource>Controller`
  - Razor Pages: `<PageName>Model` (standard convention)
  - Repository interfaces: `I<Entity>Repository`
  - DTOs: `<Entity>Request` / `<Entity>Response`

### keepee (Go)

- **Package names:** short, lowercase, no underscores (e.g., `collector`, `api`, `config`).
- **Error wrapping:** Always wrap errors with context: `fmt.Errorf("collecting packages: %w", err)`.
- **Logging:** Use `log/slog` with structured key-value pairs.
- **Configuration:** Use a `Config` struct populated from environment variables and/or CLI flags (no YAML files for the agent).
- **HTTP client:** Reuse a single `http.Client` instance; set explicit timeouts.
- **Naming:**
  - Exported types for data models: `AgentRegistration`, `SystemSnapshot`
  - Collector functions: `CollectPackages()`, `CollectOSInfo()`, `CollectDockerState()`

---

## Architecture Patterns

- The `keepee` agent is **always the initiating party** — it pushes data to `keeper` via REST.
- Authentication is **API-key per agent** (`Authorization: Bearer <key>` header).
- Data is stored as **JSONB snapshots** in PostgreSQL; avoid rigid schema migrations for new telemetry fields.
- The `keeper` never opens connections to agents (pull model is via embedded instructions in ping responses).

---

## Commit and PR Conventions

- **Conventional Commits**: `feat(keepee): ...`, `fix(keeper): ...`, `docs: ...`
- Scopes: `keeper`, `keepee`, `docs`, `ci`, `api`
- Breaking changes: append `!` or use `BREAKING CHANGE:` footer

---

## What to Avoid

- Avoid introducing new external dependencies without discussion.
- Avoid full ORM (e.g., Entity Framework) in `keeper` — use Dapper (a lightweight micro-ORM) and raw SQL only.
- Avoid global state in `keepee`.
- Avoid hard-coding the keeper URL or API keys anywhere.
- Avoid C# nullable warnings — enable `<Nullable>enable</Nullable>` and handle nulls explicitly.
