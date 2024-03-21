package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

// 这里的断言操作
func getCurrentUserID(c *gin.Context) (userId uint64, err error) {
	_userId, ok := c.Get("userID")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = _userId.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
