package logic

import (
	"forum/dao/database"
	"forum/model"
)

func GetCommunity() (data []model.Community, err error) {
	return database.GetCommunityList()
}

func GetCommunityDetail(id uint64) (community *model.CommunityDetail, err error) {
	return database.GetCommunityDetailByID(id)
}
