package database

import "forum/model"

func CreatePost(post *model.Post) (err error) {
	err = db.Create(post).Error
	return
}
