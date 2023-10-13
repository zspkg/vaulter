package vaulter

import (
	vault "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"os"
)

var ErrInvalidEnvs = errors.New("environment variables are not configured properly")

func tokenAuthCliFromEnv() (Vaulter, error) {
	var (
		vaultPath        = os.Getenv(EnvVaultPath)
		vaultToken       = os.Getenv(EnvVaultToken)
		vaultKvMountPath = os.Getenv(EnvVaultKVMountPath)
	)
	if vaultPath == "" || vaultToken == "" || vaultKvMountPath == "" {
		return nil, errors.Wrap(ErrInvalidEnvs, "invalid envs", logan.F{
			"vault_path":          vaultPath,
			"vault_token":         vaultToken,
			"vault_kv_mount_path": vaultKvMountPath,
		})

	}

	cfg := vault.DefaultConfig()
	cfg.Address = vaultPath

	vaultClient, err := vault.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault vaultClient")
	}

	vaultClient.SetToken(vaultToken)

	return &vaulter{vaultClient, vaultKvMountPath}, nil
}

func MustFromEnv(authType AuthType) Vaulter {
	var (
		err    error
		client Vaulter
	)

	switch authType {
	case AuthTypeToken:
		if client, err = tokenAuthCliFromEnv(); err == nil {
			return client
		}
	case AuthTypeCertificate:
		panic("not implemented")
	default:
		panic("unknown auth type")
	}

	panic(errors.Wrap(err, "failed to set up vaulter from environment"))
}

// TODO: IMPLEMENT ME
func certificateAuthCliFromEnv() (Vaulter, error) {
	vaultPath, vaultKvMountPath := os.Getenv(EnvVaultPath), os.Getenv(EnvVaultKVMountPath)
	if vaultPath == "" || vaultKvMountPath == "" {
		return nil, errors.Wrap(ErrInvalidEnvs, "failed to read envs", logan.F{
			"vault_path":          vaultPath,
			"vault_kv_mount_path": vaultKvMountPath,
		})
	}

	cfg := vault.DefaultConfig()
	cfg.Address = cfg.Address

	// configuring tls
	if err := cfg.ConfigureTLS(&vault.TLSConfig{
		CACert:        "",
		CACertBytes:   nil,
		CAPath:        "",
		ClientCert:    "",
		ClientKey:     "",
		TLSServerName: "",
		Insecure:      false,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to configure TLS")
	}

	// creating client based on cfg
	vaultClient, err := vault.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault vaultClient")
	}

	// authorizing vault using certificates and obtaining auth token
	response, err := vaultClient.Logical().Write("auth/cert/login", map[string]interface{}{
		"name":        "certificate",
		"common_name": "your-common-name",
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize vault client")
	}

	vaultClient.SetToken(response.Auth.ClientToken)
	return &vaulter{vaultClient, vaultKvMountPath}, nil
}
