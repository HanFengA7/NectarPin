package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var err error
var AppConfig *Config
var DB *gorm.DB
var Router *gin.Engine

func Init() {
	// [1]加载配置文件
	if AppConfig, err = LoadConfig("config.yaml"); err != nil {
		panic(err)
	}

	// [2]初始化数据库
	if DB, err = InitDB(AppConfig); err != nil {
		panic(err)
	}

	// [3]初始化路由
	if Router, err = InitRouter(AppConfig.Server.Port); err != nil {
		panic(err)
	}
}
