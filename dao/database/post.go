package database

import (
	"database/sql"
	"forum/model"
	"go.uber.org/zap"
)

func CreatePost(post *model.Post) (err error) {
	err = db.Create(post).Error
	return
}

func GetPostByID(postID string) (post *model.Post, err error) {
	post = new(model.Post)
	err = db.Where("post_id = ?", postID).Find(post).Error
	return
}

func GetPostList(page, size int) ([]model.Post, error) {
	var postList []model.Post
	err := db.Offset((page - 1) * size).Limit(size).Find(&postList).Error

	if err == sql.ErrNoRows {
		zap.L().Warn("There is no post record in DB.")
	}
	return postList, nil
}
