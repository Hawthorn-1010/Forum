package database

import (
	"forum/model"
)

func CreateUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func QueryUser(user *model.User) (dbUser *model.User, err error) {
	// add or "error": "unsupported destination, should be slice or struct"
	dbUser = new(model.User)
	err = db.Where("username = ?", user.Username).First(dbUser).Error
	return
}

func CheckUserExist(username string) error {
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return ErrorUserExit
}
