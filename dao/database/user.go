package database

import (
	"forum/model"
)

func CreateUser(user *model.User) (err error) {
	err = db.Create(user).Error
	return
}

func QueryUserByUsername(username string) (user *model.User, err error) {
	// add or "error": "unsupported destination, should be slice or struct"
	user = new(model.User)
	err = db.Where("username = ?", username).First(user).Error
	return
}

func QueryUserByID(userId string) (user *model.User, err error) {
	user = new(model.User)
	err = db.Where("user_id = ?", userId).First(user).Error
	return
}

func CheckUserExist(username string) error {
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return ErrorUserExit
}
