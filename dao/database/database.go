package database

import (
	"fmt"
	"forum/setting"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

var db *gorm.DB

func Init(cfg *setting.DBConfig) (err error) {
	driverName := cfg.DbType
	loc := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		url.QueryEscape(loc))
	db, err = gorm.Open(driverName, dsn)
	if err != nil {
		return err
	}

	// 自动创建数据表
	//db.AutoMigrate(&model.User{})
	return nil
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
