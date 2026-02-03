package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//success response

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    2001,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, massage string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    2001,
		Message: msg[code],
		Data:    nil,
	})
}
