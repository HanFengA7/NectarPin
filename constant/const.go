package constant

import (
	"NectarPin/conf"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)
