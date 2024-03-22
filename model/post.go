package model

import "time"

type Post struct {
	PostID      uint64    `json:"post_id,string" gorm:"column:post_id"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content" gorm:"column:content"`
	AuthorId    uint64    `json:"author_id" gorm:"column:author_id"`
	CommunityID uint64    `json:"community_id" gorm:"column:community_id"`
	Status      int32     `json:"status" gorm:"column:status"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}

//func (p *Post) UnmarshalJSON(data []byte) (err error) {
//	required := struct {
//		Title       string `json:"title" gorm:"column:title"`
//		Content     string `json:"content" gorm:"column:content"`
//		CommunityID int64  `json:"community_id" gorm:"column:community_id"`
//	}{}
//	err = json.Unmarshal(data, &required)
//	if err != nil {
//		return
//	} else if len(required.Title) == 0 {
//		err = errors.New("帖子标题不能为空")
//	} else if len(required.Content) == 0 {
//		err = errors.New("帖子内容不能为空")
//	} else if required.CommunityID == 0 {
//		err = errors.New("未指定版块")
//	} else {
//		p.Title = required.Title
//		p.Content = required.Content
//		p.CommunityID = required.CommunityID
//	}
//	return
//}

type ApiPostDetail struct {
	*Post
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}
