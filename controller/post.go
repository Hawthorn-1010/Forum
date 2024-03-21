package controller

import (
	"forum/logic"
	"forum/model"
	"forum/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBind(&post); err != nil {
		zap.L().Error("Create Post with invalid params", zap.Error(err))
		response.Fail(c, response.CodeInvalidParams)
		return
	}

	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		response.Fail(c, response.CodeNotLogin)
		return
	}
	post.AuthorId = userID

	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Uint64("post_id", post.PostID), zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	response.Success(c, nil)
}

func PostDetailHandler(c *gin.Context) {
	id := c.Param("id")
	postDetail, err := logic.GetPost(id)
	if err != nil {
		zap.L().Error("logic.GetPost() failed", zap.String("post_id", id), zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	response.Success(c, postDetail)
}

func PostListHandler(c *gin.Context) {

}

func PostList2Handler(c *gin.Context) {

}
