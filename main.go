package vaulter

import (
	vault "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/kit/kv"
)

const (
	// Auth types

	// AuthTypeToken is used for vault client configuration via VAULT_TOKEN environment variable
	AuthTypeToken AuthType = iota
	// AuthTypeCertificate is for another type of vault client configuration (in development)
	AuthTypeCertificate

	// ENV variables

	EnvVaultPath        = "VAULT_PATH"
	EnvVaultToken       = vault.EnvVaultToken
	EnvVaultKVMountPath = "VAULT_KV_MOUNT_PATH"
)

type Vaulter interface{ kv.Getter }

type AuthType int
