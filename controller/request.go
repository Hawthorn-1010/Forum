package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
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

func getPageInfo(c *gin.Context) (int, int) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return int(page), int(size)
}
