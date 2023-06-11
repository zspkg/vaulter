# Vaulter
[![Go Reference](https://pkg.go.dev/badge/github.com/zspkg/vaulter.svg)](https://pkg.go.dev/github.com/zspkg/vaulter)
[![Go Report Card](https://goreportcard.com/badge/github.com/zspkg/vaulter)](https://goreportcard.com/report/github.com/zspkg/vaulter)

Simple HashiCorp vault client wrapper for configuring services

## Usage example

```go
package example

import "github.com/zspkg/vaulter"

const fooVaultSecretKey = "key"

type FooConfig struct {
	Foo string `fig:"foo_key,required"`
	Bar int    `fig:"bar_secret,required"`
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
```
