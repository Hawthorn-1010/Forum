package controller

import (
	"errors"
	"forum/dao/database"
	"forum/logic"
	"forum/response"
	"forum/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var signUpParams vo.SignUpParams

	// 可根据SignUpParams的标签进行数据验证
	if err := c.ShouldBind(&signUpParams); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		response.Fail(c, response.CodeInvalidParams)
		return
	}

	// 注册用户
	err := logic.SignUp(signUpParams)

	response.Success(c, err)
}

func LoginHandler(c *gin.Context) {
	var loginParams vo.LoginParams

	if err := c.ShouldBind(&loginParams); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		response.Fail(c, response.CodeInvalidParams)
		return
	}

	token, err := logic.Login(&loginParams)

	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", loginParams.Username), zap.Error(err))
		if errors.Is(err, database.ErrorUserNotExit) {
			response.Fail(c, response.CodeUserNotExist)
			return
		}
		response.Fail(c, response.CodeInvalidPassword)
		return
	}

	response.Success(c, token)
}

func RefreshTokenHandler(c *gin.Context) {

}
