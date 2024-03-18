package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//{
//	code : 20001,
//	msg : xx
//	data : xx,
//}

type Response struct {
	Code ResCode
	Msg  string
	Data interface{}
}

func Success(c *gin.Context, data interface{}) {
	res := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func Fail(c *gin.Context, errCode ResCode) {
	res := &Response{
		Code: errCode,
		Msg:  errCode.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}
