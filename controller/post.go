package controller

import (
	"forum/logic"
	"forum/model"
	"forum/response"
	"forum/vo"
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
	page, size := getPageInfo(c)
	posts, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	response.Success(c, posts)
}

// GetPostListHandler2 根据前端传来的参数动态获取帖子列表接口
// 按创建时间或者点赞分数
func GetPostListHandler2(c *gin.Context) {
	//获取flag（获取时间排序的帖子还是点赞分数）
	//初始化结构体指定初始参数
	p := vo.ParamPostList{
		Page:  1,
		Size:  10,
		Order: vo.OrderTime,
	}
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		response.Fail(c, response.CodeInvalidParams)
		return
	}
	//获取帖子的数据
	data, err := logic.GetPostList2(&p)
	if err != nil {
		zap.L().Error("logic.GetPostList2", zap.Error(err))
		response.Fail(c, response.CodeServerBusy)
		return
	}
	//返回信息
	response.Success(c, data)

}
