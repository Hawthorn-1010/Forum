package vo

// ParamVoteData 投票参数
type ParamVoteData struct {
	// 谁(userID)给哪个帖子(postId)投了什么票(direction)
	//UserID 从请求中获取当前用户
	PostID    int64 `json:"post_id,string" binding:"required"`        //帖子id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1" ` //赞成票（1）反对票（-1）取消投票(0)
}
