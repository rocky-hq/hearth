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

## Cross-language types (Phase 5b onward)

After 5b lands, types referenced by the `Driver` interface come from `github.com/rocky-hq/contracts/go/hearth` (generated from the canonical zod source in `contracts/src/hearth/`). Until then, `internal/driver/types_local.go` carries placeholder definitions marked `// PHASE-5A-LOCAL`. **Do not extend `types_local.go` — extend the zod source instead.**
