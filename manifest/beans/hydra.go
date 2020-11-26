package beans

import (
	"github.com/go-openapi/strfmt"
	"github.com/mix-go/bean"
	SDK "github.com/ory/hydra/sdk/go/hydra/client"
	"hydra-wework-auth-server/globals"
)

func HydraClientConfig() {
	Beans = append(Beans,
		bean.BeanDefinition{
			Name:    "hydraClient",
			Reflect: bean.NewReflect(SDK.NewHTTPClientWithConfig),
			Scope:   bean.SINGLETON,
			ConstructorArgs: bean.ConstructorArgs{
				strfmt.Default,
				&SDK.TransportConfig{
					Schemes:  []string{globals.Config.HydraConfig.HydraMethod},
					BasePath: "/",
					Host:     globals.Config.HydraConfig.HydraHost + ":" + globals.Config.HydraConfig.HydraAdminPort,
				},
			},
		},
	)
}
