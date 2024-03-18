package controller

import (
	"forum/logic"
	"forum/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var signUpParams vo.SignUpParams

	// 可根据SignUpParams的标签进行数据验证
	if err := c.ShouldBind(&signUpParams); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		Fail(c, CodeInvalidParams)
		return
	}

	// 注册用户
	err := logic.SignUp(signUpParams)

	Success(c, err)
}

func LoginHandler(c *gin.Context) {

}
