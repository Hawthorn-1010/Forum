package model

import "time"

type Community struct {
	CommunityID   uint64 `json:"community_id" gorm:"column:community_id"`
	CommunityName string `json:"community_name" gorm:"column:community_name"`
}

type CommunityDetail struct {
	CommunityID   uint64 `json:"community_id" gorm:"column:community_id"`
	CommunityName string `json:"community_name" gorm:"column:community_name"`
	//omitempty表示如果为空的话就不展示
	Introduction string `json:"introduction,omitempty" gorm:"column:introduction"`
	//使用time.time类型的时候,数据库用的是时间戳,所以连接mysql要加上parseTime=true
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
