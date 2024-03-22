package controller

import (
	"forum/logic"
	"forum/response"
	"forum/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 投票功能
func VoteHandler(c *gin.Context) {
	//参数校验
	p := new(vo.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("请求参数失败", zap.Error(err))
		response.Fail(c, response.CodeInvalidParams)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		response.Fail(c, response.CodeNotLogin)
		return
	}
	//业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	//返回响应
	response.Success(c, nil)
}
