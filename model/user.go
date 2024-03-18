package model

type User struct {
	UserID   uint64 `json:"user_id" gorm:"column:user_id""`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}
