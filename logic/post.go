package logic

import (
	"fmt"
	"forum/dao/database"
	"forum/model"
	"go.uber.org/zap"
)

func CreatePost(post *model.Post) (err error) {
	return database.CreatePost(post)
}

func GetPost(postID string) (post *model.ApiPostDetail, err error) {
	post, err = database.GetPostByID(postID)
	if err != nil {
		zap.L().Error("database.GetPostByID(postID) failed", zap.String("post_id", postID), zap.Error(err))
		return nil, err
	}
	user, err := database.QueryUserByID(fmt.Sprint(post.AuthorId))
	if err != nil {
		zap.L().Error("database.QueryUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		return
	}
	post.AuthorName = user.Username
	community, err := database.GetCommunityByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("database.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	post.CommunityName = community.CommunityName
	return post, nil
}

func GetPostList2() (data []*model.ApiPostDetail, err error) {
	postList, err := database.GetPostList()
	if err != nil {
		fmt.Println(err)
		return
	}
	data = make([]*model.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		user, err := database.QueryUser(fmt.Sprint(post.AuthorId))
		if err != nil {
			zap.L().Error("database.QueryUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
			continue
		}
		post.AuthorName = user.Username
		community, err := database.GetCommunityByID(fmt.Sprint(post.CommunityID))
		if err != nil {
			zap.L().Error("database.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
			continue
		}
		post.CommunityName = community.CommunityName
		data = append(data, post)
	}
	return
}
