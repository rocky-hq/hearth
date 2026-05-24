package driver

import "context"

// Driver is the locked HEARTH provisioning protocol surface.
//
// Phase 5 spec §5 contract guarantees (enforced by
// internal/driver/fake/fake_test.go and reused by every future
// concrete driver):
//
//   - Provision is idempotent on (slug, profile.Tier, profile.Driver):
//     calling it twice returns the same DeploymentRef with no side
//     effects.
//   - Status is read-only. Callable any number of times; never mutates
//     state.
//   - Upgrade may mutate the live deployment but MUST preserve
//     DeploymentRef.WorkspaceSlug. The returned ref may change
//     Endpoint, SecretsVaultPath, LastStatus.
//   - Teardown is irreversible. After it returns, Status against the
//     same ref returns StatusTierTornDown (a terminal state, not an
//     error).
//   - All four methods MUST honour ctx.Done() and return promptly with
//     ctx.Err() when cancelled.
type Driver interface {
	Provision(ctx context.Context, slug string, profile ProvisioningProfile) (DeploymentRef, error)
	Status(ctx context.Context, ref DeploymentRef) (Status, error)
	Upgrade(ctx context.Context, ref DeploymentRef, profile ProvisioningProfile) (DeploymentRef, error)
	Teardown(ctx context.Context, ref DeploymentRef) error
}
