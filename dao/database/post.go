package database

import "forum/model"

func CreatePost(post *model.Post) (err error) {
	err = db.Create(post).Error
	return
}

func GetPostByID(postID string) (post *model.Post, err error) {
	post = new(model.Post)
	err = db.Where("post_id = ", postID).Find(post).Error
	return
}
