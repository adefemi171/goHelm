package route

import (
	controller "github.com/adefemi171/goHelm/pkg/controller"

	"github.com/gin-gonic/gin"
)

// UserRoutes for User Details and Header
func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.Welcome)
	router.GET("/users", controller.UserHandler)
	router.GET("/header", controller.GetHeader)
	router.NoRoute(controller.NotFound)
}
