package types

// ProviderName defines a type for RPC provider names.
type ProviderName string

const (
	// ProviderDefault represents the default RPC provider set.
	ProviderDefault ProviderName = "default"
	// ProviderPublic represents the public RPC provider set.
	ProviderPublic ProviderName = "public"
)
