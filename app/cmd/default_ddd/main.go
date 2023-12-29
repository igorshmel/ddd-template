package main

import (
	"default_ddd/app/cmd/default_ddd/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(bundlefx.Module).Run()
}
