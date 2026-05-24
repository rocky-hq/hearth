// Package driver declares the HEARTH provisioning surface.
//
// The Driver interface is the locked Phase 5 protocol surface
// (see rocky-hq/docs/specs/2026-05-04-rocky-phase-5.md §5). Every
// concrete driver — FakeDriver (Phase 5a), LocalDocker (Phase 5c),
// Kustomize (Phase 6a), DevarnoCloud (Phase 6b) — must satisfy it.
//
// Types referenced by the interface live in this package as
// PHASE-5A-LOCAL placeholders in types_local.go. Phase 5b replaces
// them with imports from github.com/rocky-hq/contracts/go/hearth.
package driver
