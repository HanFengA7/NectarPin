package constant

import (
	"NectarPin/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
