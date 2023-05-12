package vaulter

import (
	"context"
	vault "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	// ENV variables
	vaultPathEnv  = "VAULT_PATH"
	vaultTokenEnv = "VAULT_TOKEN"

	// Path variables
	vaultMountPath = "secret"
)

var ErrNoVault = errors.New("no vault configured")

// vaulter is a HashiCorp vault wrapper that
// implements Vaulter interface
type vaulter struct {
	vaultClient *vault.Client
}

func (c *vaulter) GetVaultSecret(key string, out any, hooks figure.Hooks) error {
	data, err := c.GetVaultSecretData(key)
	if err != nil {
		return errors.Wrap(err, "failed to get secret")
	}

	return figure.
		Out(out).
		With(hooks).
		From(data).
		Please()
}

func (c *vaulter) GetVaultSecretData(key string) (map[string]interface{}, error) {
	secret, err := c.vaultClient.KVv2(vaultMountPath).Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	return secret.Data, nil
}
