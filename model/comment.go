package model

import "time"

type Comment struct {
	PostID     uint64    `gorm:"column:question_id" json:"question_id"`
	ParentID   uint64    `gorm:"column:parent_id" json:"parent_id"`
	CommentID  uint64    `gorm:"column:comment_id" json:"comment_id"`
	AuthorID   uint64    `gorm:"column:author_id" json:"author_id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}
