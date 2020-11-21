/*
  Server Router
*/
package routers

import (
	"app/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleNotFound(c *gin.Context) {
	response.Response(c,
		http.StatusNotFound,
		http.StatusNotFound,
		fmt.Sprintf("Route '%v %v' Not Found", c.Request.Method, c.Request.URL.Path),
		map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})
}

func handleCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func Init(router *gin.Engine) {
	router.NoRoute(handleNotFound)
	router.NoMethod(handleNotFound)

	router.Use(handleCors())

	Routers(router)
}
