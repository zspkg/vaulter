# vaulter
Simple `HashiCorp` vault client wrapper for configs that is compatible with `gitlab.com/distributed_lab/kit/kv` package

## Usage example

```go
package example

import (
	"github.com/zspkg/vaulter"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

const fooVaultSecretKey = "key"

type FooConfig struct {
	Foo string `fig:"foo_key,required"`
	Bar int    `fig:"bar_secret,required"`
}

func GetFooConfig() (FooConfig, error) {
	var (
		cfg         FooConfig
		vaultGetter = vaulter.MustFromEnv(vaulter.AuthTypeToken)
	)

	fooMap := kv.MustGetStringMap(vaultGetter, fooVaultSecretKey)

	// TODO: figure out config from fooMap with your favorite config library
	if err := figure.Out(&cfg).From(fooMap).Please(); err != nil {
		return cfg, err
	}

	return cfg, nil
}
```