/*
  Response definition
*/
package response

import (
	"net/http"

	"app/status"

	"github.com/gin-gonic/gin"
)

// Body : empty for default value
type Body struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// Response is low level gin json response
func Response(ctx *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	ctx.JSON(httpStatus, &Body{Code: code, Msg: msg, Data: data})
}

// Return is return api result
func Return(ctx *gin.Context, code status.Code, data interface{}) {
	Response(ctx, http.StatusOK, int(code), code.String(), data)
}
