package routers

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

// Add routers here
func Routers(router *gin.Engine) {
	router.GET("/", controllers.Index)
}
