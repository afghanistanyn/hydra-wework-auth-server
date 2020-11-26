package manifest

import (
	"github.com/mix-go/console"
	"github.com/mix-go/dotenv"
	"hydra-wework-auth-server/manifest/beans"
	"hydra-wework-auth-server/manifest/commands"
)

var (
	ApplicationDefinition console.ApplicationDefinition
)

func Init() {
	beans.Init()

	ApplicationDefinition = console.ApplicationDefinition{
		Name:     "app",
		Version:  "1.0.0-alpha",
		Debug:    dotenv.Getenv("APP_DEBUG").Bool(),
		Beans:    beans.Beans,
		Commands: commands.Commands,
	}
}
