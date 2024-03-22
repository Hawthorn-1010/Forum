package logic

import (
	"fmt"
	"forum/dao/database"
	"forum/dao/redis"
	"forum/model"
	"forum/pkg/snowflake"
	"forum/vo"
	"go.uber.org/zap"
)

func CreatePost(post *model.Post) (err error) {
	post.PostID, _ = snowflake.GetID()

	err = database.CreatePost(post)
	if err != nil {
		return err
	}
	err = redis.CreatePost(post.PostID)
	return
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

func GetPostList(page, size int) (postDetails []*model.ApiPostDetail, err error) {
	posts, err := database.GetPostList(page, size)
	if err != nil {
		zap.L().Error("database.GetPostList() failed", zap.Error(err))
	}

	for _, post := range posts {
		postDetail, err := GetPost(fmt.Sprint(post.PostID))
		if err != nil {
			zap.L().Error("database.GetPost failed", zap.Uint64("AuthorID", post.PostID), zap.Error(err))
			continue
		}
		postDetails = append(postDetails, postDetail)
	}

	return
}

// 按照
func GetPostList2(p *vo.ParamPostList) (postDetails []*model.ApiPostDetail, err error) {
	//redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn(" redis.GetPostIDsInOrder get ids return 0 data")
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	//根据id查询post列表数据
	//按照给定的顺序返回
	posts, err := database.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("database.GetPostListByIDs", zap.Any("posts", posts))
	//根据post获取作者和社区信息填充到帖子中
	for _, post := range posts {
		postDetail, err := GetPost(fmt.Sprint(post.PostID))
		if err != nil {
			zap.L().Error("database.GetPost failed", zap.Uint64("AuthorID", post.PostID), zap.Error(err))
			continue
		}
		postDetails = append(postDetails, postDetail)
	}

	return
}
