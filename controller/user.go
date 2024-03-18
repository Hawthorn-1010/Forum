package controller

import (
	"forum/response"
	"forum/vo"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var signUpParams vo.SignUpParams

	// 可根据SignUpParams的标签进行数据验证
	if err := c.ShouldBind(&signUpParams); err != nil {
		response.Fail(c, response.CodeInvalidParams)
		return
	}

	// 注册用户
}

func LoginHandler(c *gin.Context) {

}
