package vaulter

import (
	vault "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"os"
)

func tokenFromEnv() (Vaulter, error) {
	vaultPath, vaultToken := os.Getenv(vaultPathEnv), os.Getenv(vaultTokenEnv)
	if vaultPath == "" || vaultToken == "" {
		return nil, ErrNoVault
	}

	cfg := vault.DefaultConfig()
	cfg.Address = vaultPath

	vaultClient, err := vault.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault vaultClient")
	}

	vaultClient.SetToken(vaultToken)

	return &vaulter{vaultClient}, nil
}

func MustFromEnv(authType AuthType) Vaulter {
	switch authType {
	case AuthTypeToken:
		client, err := tokenFromEnv()
		if err != nil {
			panic(errors.Wrap(err, "failed to set up vaulter vaulter from environment"))
		}

		return client
	case AuthTypeCertificate:
		//TODO: IMPLEMENT ME
		panic("Not implemented yet")
	default:
		panic("unknown auth type")
	}
}
