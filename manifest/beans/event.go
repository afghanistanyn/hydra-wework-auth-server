package beans

import (
	"github.com/mix-go/bean"
	"github.com/mix-go/event"
	"hydra-wework-auth-server/listeners"
)

func Event() {
	Beans = append(Beans,
		bean.BeanDefinition{
			Name:            "eventDispatcher",
			Reflect:         bean.NewReflect(event.NewDispatcher),
			Scope:           bean.SINGLETON,
			ConstructorArgs: bean.ConstructorArgs{&listeners.CommandListener{}},
		},
	)
}
