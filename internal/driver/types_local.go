package driver

// PHASE-5A-LOCAL — these types will be replaced wholesale in Phase 5b
// by imports from github.com/rocky-hq/contracts/go/hearth (generated
// from the canonical zod source). Do NOT extend them here; extend the
// zod source in contracts/src/hearth/ instead.
//
// Shapes match Phase 5 spec §6 (TierSchema, DriverNameSchema,
// ProvisioningProfileSchema, DeploymentRefSchema, StatusSchema).

// Tier is the workspace tier (TierSchema).
type Tier string

const (
	TierSolo    Tier = "solo"
	TierTeam    Tier = "team"
	TierStudio  Tier = "studio"
	TierBespoke Tier = "bespoke"
)

// DriverName identifies the driver implementation (DriverNameSchema).
type DriverName string

const (
	DriverLocalDocker  DriverName = "local-docker"
	DriverKustomize    DriverName = "kustomize"
	DriverDevarnoCloud DriverName = "devarno-cloud"
)

// Status is the deployment lifecycle state (StatusSchema).
type Status string

const (
	StatusProvisioning Status = "provisioning"
	StatusReady        Status = "ready"
	StatusUpgrading    Status = "upgrading"
	StatusTearingDown  Status = "tearing_down"
	StatusFailed       Status = "failed"
	StatusTierTornDown Status = "tier_torn_down"
)

// ResourceCaps mirrors ProvisioningProfileSchema.resource_caps.
type ResourceCaps struct {
	CairnetStorageMB  int    `json:"cairnet_storage_mb"`
	LoreRetentionDays int    `json:"lore_retention_days"`
	Seats             int    `json:"seats"`
	VectorIndex       string `json:"vector_index"` // "faiss-local" | "pgvector"
}

// ProvisioningProfile mirrors ProvisioningProfileSchema.
type ProvisioningProfile struct {
	Tier         Tier                   `json:"tier"`
	Driver       DriverName             `json:"driver"`
	ResourceCaps ResourceCaps           `json:"resource_caps"`
	DriverFlags  map[string]interface{} `json:"driver_flags"`
}

// DeploymentRef mirrors DeploymentRefSchema.
type DeploymentRef struct {
	WorkspaceSlug    string     `json:"workspace_slug"`
	Tier             Tier       `json:"tier"`
	Driver           DriverName `json:"driver"`
	Endpoint         string     `json:"endpoint"`
	SecretsVaultPath string     `json:"secrets_vault_path"`
	Created          string     `json:"created"` // ISO8601
	LastStatus       Status     `json:"last_status"`
}
