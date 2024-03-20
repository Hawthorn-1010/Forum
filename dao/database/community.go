package database

import (
	"database/sql"
	"forum/model"
	"go.uber.org/zap"
)

func GetCommunityList() ([]model.Community, error) {
	//	查询数据
	//communityList = new(model.Community)
	// TODO
	var communityList []model.Community
	err := db.Find(&communityList).Error

	if err == sql.ErrNoRows {
		zap.L().Warn("There is no community record in DB.")
	}
	return communityList, nil
}

func GetCommunityDetailByID(id uint64) (community *model.CommunityDetail, err error) {
	community = new(model.CommunityDetail)
	err = db.Where("community_id = ?", id).First(community).Error

	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.Uint64("id", id), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}
