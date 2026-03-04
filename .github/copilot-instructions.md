# Copilot Instructions for semantic-conventions-go

## Project Overview

`semantic-conventions-go` is a Go library that provides typed constants and helpers for [OpenTelemetry Semantic Conventions](https://opentelemetry.io/docs/specs/semconv/). It allows Go applications to reference well-known attribute keys, values, and namespaces without relying on hard-coded strings.

## Repository Structure

```
.
├── .github/                  # GitHub configuration (workflows, instructions)
├── LICENSE
├── README.md
└── (Go packages added under appropriate subdirectories)
```

Go packages should be organized by semantic convention namespace (e.g., `http`, `db`, `messaging`) and placed under the root module path.

## Prerequisites

- Go 1.21 or later (check `go.mod` for the exact minimum version once initialized)
- Standard Go toolchain (`go build`, `go test`, `go vet`)

## Build & Test

```bash
# Build all packages
go build ./...

# Run all tests
go test ./...

# Run tests with race detector
go test -race ./...

# Run go vet
go vet ./...
```

## Linting

This project uses [golangci-lint](https://golangci-lint.run/). Run it before submitting changes:

```bash
golangci-lint run ./...
```

If a `.golangci.yml` configuration file is present at the root, golangci-lint will pick it up automatically.

## Code Conventions

- Follow standard [Go code style](https://go.dev/doc/effective_go) and `gofmt` formatting.
- All exported symbols must have Go doc comments.
- Attribute key constants should be typed as `attribute.Key` (from `go.opentelemetry.io/otel/attribute`).
- Use `const` blocks for groups of related attribute keys.
- Keep package names lowercase and matching their semantic convention namespace.
- Do not introduce dependencies outside of the OpenTelemetry Go SDK and the Go standard library unless absolutely necessary.

## CI / Validation

Before opening a pull request, ensure:

1. `go build ./...` succeeds with no errors.
2. `go test -race ./...` passes.
3. `go vet ./...` reports no issues.
4. `golangci-lint run ./...` reports no issues.
5. All exported symbols have doc comments.

## Dependency Management

- Use Go modules (`go mod tidy` after adding or removing dependencies).
- Do not vendor dependencies (`vendor/` is not committed).
- Pin dependency versions in `go.mod`/`go.sum`.
