package controller

import (
	"errors"
	"fmt"
	"forum/dao/database"
	"forum/logic"
	"forum/pkg/jwt"
	"forum/response"
	"forum/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
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

	accessToken, refreshToken, err := logic.Login(&loginParams)

	if err != nil {
		zap.L().Error("logic.Login() failed", zap.String("username", loginParams.Username), zap.Error(err))
		if errors.Is(err, database.ErrorUserNotExit) {
			response.Fail(c, response.CodeUserNotExist)
			return
		}
		response.Fail(c, response.CodeInvalidPassword)
		return
	}

	response.Success(c, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		response.FailWithMsg(c, response.CodeInvalidToken, "请求头缺少AuthToken")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.FailWithMsg(c, response.CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	accessToken, refreshToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	response.Success(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
