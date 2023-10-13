package vaulter

import (
	"context"
	vault "github.com/hashicorp/vault/api"
	"strings"
)

// vaulter is a HashiCorp vault wrapper that implements Vaulter interface
type vaulter struct {
	vaultClient *vault.Client
	mountPath   string
}

func (c *vaulter) GetStringMap(key string) (map[string]interface{}, error) {
	secret, err := c.vaultClient.KVv2(c.mountPath).Get(context.Background(), key)
	if err != nil {
		if strings.Contains(err.Error(), vault.ErrSecretNotFound.Error()) {
			return nil, nil
		}

		return nil, err
	}

	return secret.Data, nil
}
