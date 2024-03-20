package controller

import (
	"forum/logic"
	"forum/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	communityList, err := logic.GetCommunity()
	if err != nil {
		zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	response.Success(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	//如果不是字符类型,则报类型错误
	if err != nil {
		response.Fail(c, response.CodeInvalidParams)
	}
	communityDetail, err := logic.GetCommunityDetail(uint64(id))
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	response.Success(c, communityDetail)
}
