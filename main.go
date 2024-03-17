package main

import (
	"fmt"
	"forum/dao/database"
	"forum/dao/redis"
	"forum/logger"
	"forum/route"
	"forum/setting"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置（viper配置管理）
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 2. 初始化日志（zap日志库）
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 3. 初始化MySQL连接
	if err := database.Init(setting.Conf.DBConfig); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	defer database.Close()

	// 4. 初始化Redis连接（go-redis）
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5. 注册路由
	r := gin.Default()
	r = route.CollectRoute(r)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

	// 6. 优雅关机
}
