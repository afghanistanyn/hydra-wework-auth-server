package routes

import (
	"github.com/gin-gonic/gin"
	"hydra-wework-auth-server/controllers"
)

func RouteDefinitionCallbacks() (callbacks []func(router *gin.Engine)) {
	callbacks = append(callbacks, func(router *gin.Engine) {

		router.Any("/wework/auth", func(ctx *gin.Context) {
			login := controllers.LoginController{}
			login.Login(ctx)
		})

		router.Any("/wework/consent", func(ctx *gin.Context) {
			consent := controllers.ConsentController{}
			consent.Consent(ctx)
		})

		router.Any("/wework/callback", func(ctx *gin.Context) {
			callback := controllers.CallbackController{}
			callback.Callback(ctx)
		})

		router.GET("/error", func(ctx *gin.Context) {
			hydraerr := controllers.HydraErrorController{}
			hydraerr.HydraError(ctx)
		})

	})
	return
}
