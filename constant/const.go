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

const (
	ColorBlackA  = 0
	ColorRedA    = 1
	ColorGreenA  = 2
	ColorYellowA = 3
	ColorBlueA   = 4
	ColorPurpleA = 5
	ColorYoungA  = 6
	ColorAshA    = 7

	ColorBlackB  = 0
	ColorRedB    = 1
	ColorGreenB  = 2
	ColorYellowB = 3
	ColorBlueB   = 4
	ColorPurpleB = 5
	ColorYoungB  = 6
	ColorAshB    = 7
)
