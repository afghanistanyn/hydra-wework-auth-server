package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/mix-go/console"
	"github.com/mix-go/console/argv"
	"hydra-wework-auth-server/globals"
	"hydra-wework-auth-server/manifest"
)

func init() {
	// Conf support YAML, JSON, TOML, Shell Environment
	if err := configor.Load(&globals.Config, fmt.Sprintf("%s/../conf/config.json", argv.Program().Dir)); err != nil {
		panic(err)
	}
	// Manifest
	manifest.Init()
}

func main() {
	// App
	console.NewApplication(manifest.ApplicationDefinition, "eventDispatcher", "error").Run()
}
