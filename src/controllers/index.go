package controllers

import (
	"app/response"
	"app/status"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Return(c, status.OK, map[string]interface{}{
		"status": "It's work",
	})
}
