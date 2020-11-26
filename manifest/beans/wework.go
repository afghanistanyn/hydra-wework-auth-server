package beans

import (
	"github.com/mix-go/bean"
	"hydra-wework-auth-server/globals"
	"hydra-wework-auth-server/wework"
)

func WeWrorkClientConfig() {
	Beans = append(Beans,
		bean.BeanDefinition{
			Name:    "weWorkClient",
			Reflect: bean.NewReflect(wework.NewClient),
			Scope:   bean.SINGLETON,
			ConstructorArgs: bean.ConstructorArgs{
				globals.Config.WeworkConfig.WeworkCropID,
				globals.Config.WeworkConfig.WeworkAgentID,
				globals.Config.WeworkConfig.WeworkSecret,
			},
		},
	)
}
