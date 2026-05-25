# CLAUDE.md — hearth/

Per the parent superproject's push-down policy (rocky-hq decision 2026-05-02-redesign-bootstrap §3), per-language build/test/lint instructions for HEARTH live here, not in parent `CLAUDE.md`.

## What hearth is

The Go submodule implementing the per-workspace CAIRNET+LORE provisioner from `rocky-hq/docs/specs/2026-05-02-rocky-system-redesign.md` §SS-08. Surface: the four-method `Driver` interface in `internal/driver/driver.go`. Phase 5a ships the interface + `FakeDriver`; Phase 5c ships `LocalDocker`; Phases 6a/6b ship `Kustomize` and `DevarnoCloud`.

## Build

```bash
go build ./...
```

## Test

Unit (no Docker required):

```bash
go test ./... -race -count=1
```

Integration (requires Docker daemon; lands in Phase 5c):

```bash
ROCKY_HEARTH_INTEGRATION=1 go test ./test/integration/... -race -count=1
```

## Lint

```bash
gofumpt -l .              # exits clean (no diff)
go vet ./...
golangci-lint run         # uses .golangci.yml
go mod tidy && git diff --exit-code go.mod go.sum
```

## Conventions

Locked in `rocky-hq/docs/decisions/2026-05-03-phase-5-go-conventions.md`:
- Module: `github.com/rocky-hq/hearth`
- License: MIT
- Go: 1.25
- Layout: `cmd/` (entrypoints — Phase 5c+), `internal/` (private packages), `test/integration/` (gated)
- Formatter: `gofumpt`
- Linter: `golangci-lint` with the baseline `.golangci.yml`
- No vendoring, no `pkg/` directory

## Driver wire-format types

The Phase 5 spec §6 shapes (`Tier`, `DriverName`, `Status`, `ResourceCaps`,
`ProvisioningProfile`, `DeploymentRef`, `HearthHatchEvent`) are NOT
defined in this repo. They are generated from the zod source in
`@rocky-hq/contracts/src/hearth/` and imported as Go bindings from
`github.com/rocky-hq/contracts/go/hearth`. The re-exports under
`package driver` (`internal/driver/types.go`) are a thin convenience layer.
`internal/driver/types_local.go` was removed in Phase 5b — it no longer exists.

**To add or modify a wire-format type:**

1. Edit the zod schema in `contracts/src/hearth/`.
2. Bump the contracts package version and tag (both `vX.Y.Z` and `go/vX.Y.Z`).
3. `go get github.com/rocky-hq/contracts/go@<new-tag>` here.

Never edit `internal/driver/types.go` to introduce a new shape.

## CI secret: `ROCKY_HQ_RO_TOKEN`

`rocky-hq/contracts` is a private repo. The `lint`, `test`, and `integration`
CI jobs in `.github/workflows/ci.yml` configure `GOPRIVATE=github.com/rocky-hq/*`
and authenticate `go mod download` using a fine-grained PAT stored as the
org-level GitHub Actions secret **`ROCKY_HQ_RO_TOKEN`**.

If you are setting up a fresh CI runner or forking this repo:

- Create a fine-grained PAT with **read-only Contents** scope on
  `rocky-hq/contracts` (and any future private rocky-hq Go modules hearth
  might depend on).
- Store it as an org-level (or repo-level) GitHub Actions secret named
  `ROCKY_HQ_RO_TOKEN`.
- Without it, CI fails at `go mod download` with a 404 from the Go module proxy.
- Public-fork PRs cannot access org secrets and will fail CI — this tradeoff
  is documented in `.github/workflows/ci.yml` and is accepted.
