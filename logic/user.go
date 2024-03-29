package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"forum/dao/database"
	"forum/model"
	"forum/pkg/jwt"
	"forum/pkg/snowflake"
	"forum/vo"
	"go.uber.org/zap"
)

const SECRET = "MyForum"

func encryptPassword(oPassword string) string {
	h := md5.New()
	//加盐的字符串
	h.Write([]byte(SECRET))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

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
		Password: encryptPassword(params.Password),
	}
	return database.CreateUser(user)
}

func Login(params *vo.LoginParams) (accessToken, refreshToken string, err error) {
	user, err := database.QueryUserByUsername(params.Username)

	if err != nil {
		zap.L().Error("database.QueryUser() failed", zap.String("username", fmt.Sprint(params.Username)), zap.Error(err))
		return "", "", err
	}
	if user == nil {
		err = database.ErrorUserNotExit
		return
	}
	if user.Password != encryptPassword(params.Password) {
		err = database.ErrorPasswordWrong
		return
	}

	return jwt.GenToken(user.UserID, user.Username)
}
