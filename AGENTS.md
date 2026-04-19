# AGENTS.md — syskeeper

This file provides guidance for AI coding agents (e.g. OpenAI Codex, GitHub Copilot
Workspace) working in this repository. Read it before making any changes.

---

## Repository Layout

```
syskeeper/
├── keeper/          # ASP.NET Core 10 server (C#)
│   └── ...          # Razor Pages, REST API, Dapper repositories
├── keepee/          # Go 1.26+ agent binary
│   └── ...          # collectors, api client, config
├── docs/
│   ├── ARCHITECTURE.md
│   └── ADR/         # Architecture Decision Records
├── openapi/         # OpenAPI specification
├── .github/
│   ├── copilot-instructions.md
│   ├── ISSUE_TEMPLATE/
│   └── PULL_REQUEST_TEMPLATE.md
├── AGENTS.md        # this file
├── AI_GUIDELINES.md
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
└── README.md
```

---

## Build & Test Commands

### keeper (C# / ASP.NET Core 10)

```bash
cd keeper

# Restore dependencies
dotnet restore

# Build
dotnet build

# Run all tests
dotnet test

# Run the server (development)
dotnet run
```

### keepee (Go 1.26+)

```bash
cd keepee

# Download dependencies
go mod download

# Build
go build ./...

# Run all tests
go test ./...

# Format and vet
gofmt -w .
go vet ./...

# Run the agent (development)
go run . --keeper-url http://localhost:5000 --api-key <key>
```

---

## Key Conventions

### keeper (C#)

- **Namespaces:** `Keeper.<Layer>` — e.g. `Keeper.Api`, `Keeper.Data`, `Keeper.Pages`
- **Async:** All I/O methods must be `async Task<T>`; never use `.Result` or `.Wait()`
- **Data access:** Dapper only — no Entity Framework; repository classes live in `Keeper.Data`
- **Validation:** `DataAnnotations` or FluentValidation on all input models
- **Error responses:** `ProblemDetails` (RFC 7807)
- **Logging:** `ILogger<T>` with message templates; no string interpolation in log calls
- **Nullability:** `<Nullable>enable</Nullable>` is set; handle all nulls explicitly

### keepee (Go)

- **Package names:** short, lowercase, no underscores — e.g. `collector`, `api`, `config`
- **Error wrapping:** `fmt.Errorf("context: %w", err)` on every error path
- **Logging:** `log/slog` with structured key-value pairs
- **Config:** `Config` struct from environment variables / CLI flags — no YAML
- **HTTP client:** single shared `http.Client` instance with explicit timeouts
- **Tests:** table-driven tests preferred

---

## Architecture Rules (do not violate)

- `keepee` is always the initiating party — it pushes data to `keeper` via REST.
  The `keeper` never opens connections to agents.
- Authentication is **API-key per agent** (`Authorization: Bearer <key>` header).
- Snapshot telemetry is stored as **JSONB** in PostgreSQL — avoid rigid schema
  migrations for new telemetry fields.
- Do not introduce full ORMs (Entity Framework). Dapper + raw SQL only.
- Do not add new external dependencies without discussion (open an issue first).
- Do not hard-code keeper URLs or API keys anywhere in source files.

---

## Commit & PR Conventions

Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/):

```
<type>(<scope>): <short summary>
```

Valid scopes: `keeper`, `keepee`, `docs`, `ci`, `api`

Examples:

```
feat(keepee): add Docker container state collection
fix(keeper): handle duplicate agent registration gracefully
docs: update architecture diagram
chore(keeper): upgrade to .NET 11
```

Breaking changes: append `!` or add `BREAKING CHANGE:` in the footer.

---

## AI Usage Policy

See [`AI_GUIDELINES.md`](AI_GUIDELINES.md) for the project's full AI usage policy.
Key points:

- You must understand and own every change you propose.
- AI-generated code must pass the same quality bar as human-written code.
- Extra care is required in security-sensitive areas (auth, validation, crypto, SQL).
- Do not generate or commit secrets, credentials, or private keys.
- Do not reproduce copyrighted content verbatim.
