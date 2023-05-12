package vaulter

import "gitlab.com/distributed_lab/figure"

type Vaulter interface {
	// GetVaultSecret retrieves secret from the vault in the given
	// config struct. Uses figure.Hooks to fetch data from string map.
	// figure.BaseHooks will be used if no hooks are provided
	GetVaultSecret(key string, out any, hooks figure.Hooks) error
	// GetVaultSecretData retrieves raw secret map from the vault by provided key
	GetVaultSecretData(key string) (map[string]interface{}, error)
}
