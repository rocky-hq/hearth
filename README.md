# HEARTH

Per-workspace CAIRNET+LORE provisioner for the [Rocky](https://github.com/rocky-hq) superproject.

> **Status:** Phase 5a — scaffold only. The `Driver` interface and `FakeDriver` are in place; the `LocalDocker` driver lands in Phase 5c.

## Quickstart

```go
import (
    "context"
    "github.com/rocky-hq/hearth/internal/driver"
    "github.com/rocky-hq/hearth/internal/driver/fake"
)

d := fake.New()
ref, err := d.Provision(context.Background(), "demo", driver.ProvisioningProfile{
    Tier: driver.TierSolo,
})
```

## Build & test

See [`CLAUDE.md`](./CLAUDE.md) — the canonical build/test/lint instructions per the parent superproject's push-down policy.

## License

MIT. See [`LICENSE`](./LICENSE).

## Related

- Parent superproject: [`rocky-hq/rocky-hq`](https://github.com/rocky-hq/rocky-hq)
- Phase 5 spec: `rocky-hq/docs/specs/2026-05-04-rocky-phase-5.md`
- Go conventions: `rocky-hq/docs/decisions/2026-05-03-phase-5-go-conventions.md`
