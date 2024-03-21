package logic

import (
	"fmt"
	"forum/dao/database"
	"forum/model"
	"forum/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *model.Post) (err error) {
	post.PostID, _ = snowflake.GetID()
	return database.CreatePost(post)
}

func GetPost(postID string) (data *model.ApiPostDetail, err error) {
	post, err := database.GetPostByID(postID)
	if err != nil {
		zap.L().Error("database.GetPostByID(postID) failed", zap.String("post_id", postID), zap.Error(err))
		return nil, err
	}
	user, err := database.QueryUserByID(fmt.Sprint(post.AuthorId))
	if err != nil {
		zap.L().Error("database.QueryUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		return
	}
	community, err := database.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("database.GetCommunityDetailByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	data = &model.ApiPostDetail{
		Post:          post,
		AuthorName:    user.Username,
		CommunityName: community.CommunityName,
	}
	return
}
