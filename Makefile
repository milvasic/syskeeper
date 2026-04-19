.PHONY: all build test lint format help

# ── defaults ─────────────────────────────────────────────────────────────────
all: build test lint

# ── keeper (.NET) ─────────────────────────────────────────────────────────────
.PHONY: keeper-build keeper-test keeper-lint keeper-format

keeper-build:
	cd keeper && dotnet build --configuration Release

keeper-test:
	cd keeper && dotnet test --configuration Release --no-build

keeper-lint:
	cd keeper && dotnet build --configuration Release
	cd keeper && dotnet csharpier --check .

keeper-format:
	cd keeper && dotnet csharpier .
	cd keeper && dotnet format

# ── keepee (Go) ───────────────────────────────────────────────────────────────
.PHONY: keepee-build keepee-test keepee-lint keepee-format

keepee-build:
	cd keepee && go build ./...

keepee-test:
	cd keepee && go test ./...

keepee-lint:
	cd keepee && go vet ./...
	cd keepee && golangci-lint run ./...

keepee-format:
	cd keepee && gofmt -w .
	cd keepee && goimports -w .

# ── combined targets ──────────────────────────────────────────────────────────
build: keeper-build keepee-build

test: keeper-test keepee-test

lint: keeper-lint keepee-lint

format: keeper-format keepee-format

# ── help ──────────────────────────────────────────────────────────────────────
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Combined targets:"
	@echo "  all            Build, test, and lint both projects (default)"
	@echo "  build          Build both keeper and keepee"
	@echo "  test           Run tests for both keeper and keepee"
	@echo "  lint           Lint both keeper and keepee"
	@echo "  format         Format both keeper and keepee"
	@echo ""
	@echo "keeper (.NET) targets:"
	@echo "  keeper-build   dotnet build"
	@echo "  keeper-test    dotnet test"
	@echo "  keeper-lint    dotnet build + csharpier --check"
	@echo "  keeper-format  csharpier + dotnet format"
	@echo ""
	@echo "keepee (Go) targets:"
	@echo "  keepee-build   go build ./..."
	@echo "  keepee-test    go test ./..."
	@echo "  keepee-lint    go vet + golangci-lint"
	@echo "  keepee-format  gofmt + goimports"
