.PHONY: build test lint generate clean help

# Build-time version injection
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

# ── build ────────────────────────────────────────────────────────────────────

build:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o keepee ./cmd/keepee

# ── test ─────────────────────────────────────────────────────────────────────

test:
	go test ./...

# ── lint ─────────────────────────────────────────────────────────────────────

lint:
	go vet ./...
	golangci-lint run ./...

# ── generate ─────────────────────────────────────────────────────────────────

generate:
	go tool oapi-codegen --config oapi-codegen.yml ../openapi/openapi.yaml

# ── clean ────────────────────────────────────────────────────────────────────

clean:
	rm -f keepee

# ── help ─────────────────────────────────────────────────────────────────────

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build      Build static binary (CGO_ENABLED=0) with version ldflags"
	@echo "  test       Run all tests"
	@echo "  lint       Run go vet and golangci-lint"
	@echo "  generate   Regenerate API client from OpenAPI spec (requires oapi-codegen)"
	@echo "  clean      Remove built binary"
