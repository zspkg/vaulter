package example

import "github.com/zspkg/vaulter"

const fooVaultSecretKey = "key"

type FooConfig struct {
	Foo string `json:"foo_key"`
	Bar int    `json:"bar_secret"`
}

func GetFooConfig() (FooConfig, error) {
	var (
		cfg         FooConfig
		vaultGetter = vaulter.MustFromEnv()
	)

	if err := vaultGetter.GetVaultSecret(fooVaultSecretKey, &cfg, nil); err != nil {
		// handle error
	}

	return cfg, nil
}
