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

//func (c CategoryRepository) Create(name string) (*model.Category, error) {
//	category := model.Category{
//		Name: name,
//	}
//	if err := c.DB.Create(&category).Error; err != nil {
//		return nil, err
//	}
//	return &category, nil
//}
//
//func (c CategoryRepository) Delete(id int) error {
//	if err := c.DB.Delete(&model.Category{}, id).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
//	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
//		return nil, err
//	}
//	return &category, nil
//}
//
//func (c CategoryRepository) Query(id int) (*model.Category, error) {
//	var category model.Category
//	if err := c.DB.First(&category, id).Error; err != nil {
//		return nil, err
//	}
//	return &category, nil
//}
