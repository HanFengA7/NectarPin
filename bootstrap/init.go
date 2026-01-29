package bootstrap

import "gorm.io/gorm"

var err error
var AppConfig *Config
var DB *gorm.DB

func Init() {
	// [1]加载配置文件
	if AppConfig, err = LoadConfig("config.yaml"); err != nil {
		panic(err)
	}

	// [2]初始化数据库
	if DB, err = InitDB(AppConfig); err != nil {
		panic(err)
	}
}