package vaulter

import "gitlab.com/distributed_lab/figure"

const (
	// Auth types

	// AuthTypeToken is used for vault client configuration via VAULT_TOKEN environment variable
	AuthTypeToken AuthType = iota
	// AuthTypeCertificate is for another type of vault client configuration (in development)
	AuthTypeCertificate

	// ENV variables

	vaultPathEnv  = "VAULT_PATH"
	vaultTokenEnv = "VAULT_TOKEN"

	// Path variables

	vaultMountPath = "secret"
)

type Vaulter interface {
	// GetVaultSecret retrieves secret from the vault in the given
	// config struct. Uses figure.Hooks to fetch data from string map.
	// figure.BaseHooks will be used if no hooks are provided
	GetVaultSecret(key string, out any, hooks figure.Hooks) error
	// GetVaultSecretData retrieves raw secret map from the vault by provided key
	GetVaultSecretData(key string) (map[string]interface{}, error)
}

type AuthType int
