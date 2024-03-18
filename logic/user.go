package logic

import (
	"forum/dao/database"
	"forum/model"
	"forum/pkg/snowflake"
	"forum/vo"
)

func SignUp(params vo.SignUpParams) error {
	// 如果数据库已经存在该用户
	if err := database.CheckUserExist(params.Username); err != nil {
		return err
	}

	id, err := snowflake.GetID()
	if err != nil {
		return database.ErrorGenIDFailed
	}
	user := &model.User{
		UserID:   id,
		Username: params.Username,
		Password: params.Password,
	}
	return database.CreateUser(user)
}
