package commands

import (
	"github.com/mix-go/console"
	"hydra-wework-auth-server/commands"
)

func init() {
	Commands = append(Commands,
		console.CommandDefinition{
			Name:  "web",
			Usage: "\tStart the api server",
			Options: []console.OptionDefinition{
				{
					Names: []string{"a", "addr"},
					Usage: "\tListen to the specified address",
				},
				{
					Names: []string{"d", "daemon"},
					Usage: "\tRun in the background",
				},
			},
			Command: &commands.WebCommand{},
		},
	)
}
