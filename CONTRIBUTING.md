# Contributing to syskeeper

Thank you for your interest in contributing to **syskeeper**! This document covers everything you need to know to get started.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
  - [Repository Layout](#repository-layout)
  - [Development Setup](#development-setup)
- [Branching Strategy](#branching-strategy)
- [Commit Conventions](#commit-conventions)
- [Pull Request Process](#pull-request-process)
- [Coding Guidelines](#coding-guidelines)
- [Running Tests](#running-tests)

---

## Code of Conduct

By participating in this project you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). Please read it before contributing.

---

## Getting Started

### Repository Layout

```
syskeeper/
├── keeper/          # ASP.NET Core server (Razor Pages + PostgreSQL + Dapper)
├── keepee/          # Go agent binary
├── docs/            # Architecture docs and ADRs
│   └── ADR/         # Architecture Decision Records
├── openapi/         # OpenAPI specification
├── .github/         # GitHub templates and Copilot instructions
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
├── LICENSE
└── README.md
```

### Development Setup

#### Prerequisites

| Tool              | Minimum version | Purpose                |
| ----------------- | --------------- | ---------------------- |
| .NET SDK          | 10.0            | Build and run `keeper` |
| Go                | 1.26            | Build and run `keepee` |
| PostgreSQL        | 15              | `keeper` database      |
| Docker (optional) | 24              | Run PostgreSQL locally |

#### Clone and configure

```bash
git clone https://github.com/milvasic/syskeeper.git
cd syskeeper
```

#### Run the keeper (server)

```bash
cd keeper
# Copy and edit the configuration
cp appsettings.json appsettings.Development.json
# Set your PostgreSQL connection string in appsettings.Development.json

dotnet restore
dotnet run
```

The dashboard will be available at `http://localhost:5000`.

#### Run the keepee (agent)

```bash
cd keepee
go mod download
go run . --keeper-url http://localhost:5000 --api-key <your-key>
```

---

## Branching Strategy

We use a trunk-based workflow with short-lived feature branches:

| Branch pattern                 | Purpose                                  |
| ------------------------------ | ---------------------------------------- |
| `main`                         | Stable, always deployable                |
| `feat/<short-description>`     | New features                             |
| `fix/<short-description>`      | Bug fixes                                |
| `chore/<short-description>`    | Maintenance, dependency updates, tooling |
| `docs/<short-description>`     | Documentation only                       |
| `refactor/<short-description>` | Refactoring without behavior change      |

**Rules:**

- Branch off `main`; target `main` in your PR.
- Delete feature branches after merging.
- Never force-push to `main`.

---

## Commit Conventions

This project follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

### Format

```
<type>(<scope>): <short summary>

[optional body]

[optional footer(s)]
```

### Types

| Type       | When to use                                             |
| ---------- | ------------------------------------------------------- |
| `feat`     | A new feature                                           |
| `fix`      | A bug fix                                               |
| `docs`     | Documentation only changes                              |
| `style`    | Formatting, missing semicolons, etc. (no logic change)  |
| `refactor` | Code change that neither fixes a bug nor adds a feature |
| `test`     | Adding or correcting tests                              |
| `chore`    | Build process, dependency updates, tooling              |
| `ci`       | CI/CD pipeline changes                                  |
| `perf`     | Performance improvements                                |
| `revert`   | Reverts a previous commit                               |

### Scope (optional)

Use the affected component: `keeper`, `keepee`, `docs`, `ci`, `api`.

### Examples

```
feat(keepee): add Docker container state collection
fix(keeper): handle duplicate agent registration gracefully
docs: add architecture diagram to README
chore(keeper): upgrade to .NET 11
```

### Breaking changes

Append `!` after the type/scope, or add `BREAKING CHANGE:` in the footer:

```
feat(api)!: rename /agents endpoint to /registrations

BREAKING CHANGE: clients must update their base paths.
```

---

## Pull Request Process

1. **Create a branch** following the naming conventions above.
2. **Make focused changes** — one logical change per PR.
3. **Write or update tests** for any functional change.
4. **Update documentation** if your change affects behaviour or public interfaces.
5. **Open a PR** against `main` using the [pull request template](.github/PULL_REQUEST_TEMPLATE.md).
6. **Ensure all CI checks pass** before requesting review.
7. **Request a review** from a maintainer. PRs require at least one approval before merging.
8. **Squash merge** is preferred to keep the history clean.

### PR Checklist

Before submitting, make sure you have:

- [ ] Added or updated tests where applicable
- [ ] Updated relevant documentation
- [ ] Run the formatter/linter for the affected component
- [ ] Verified the build passes locally
- [ ] Written a clear PR description explaining _what_ and _why_

---

## Coding Guidelines

### keeper (.NET / ASP.NET Core)

- Follow the [.NET runtime coding style guidelines](https://github.com/dotnet/runtime/blob/main/docs/coding-guidelines/coding-style.md).
- Use `async`/`await` throughout; avoid `.Result` and `.Wait()`.
- All database access goes through Dapper repository classes — no raw ADO.NET in controllers.
- Validate all inputs; never trust agent-supplied data without validation.
- Add XML doc comments (`///`) to all public APIs.

### keepee (Go)

- Follow standard Go formatting; run `gofmt` and `go vet` before committing.
- Keep the binary self-contained and dependency-light.
- Use structured logging (e.g., `log/slog`).
- Errors must be wrapped with context: `fmt.Errorf("collecting packages: %w", err)`.
- Table-driven tests are preferred.

---

## Running Tests

### keeper

```bash
cd keeper
dotnet test
```

### keepee

```bash
cd keepee
go test ./...
```
