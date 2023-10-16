package Init

import (
	"NectarPin/constant"
	"NectarPin/tools"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"time"
)

func Gorm() *gorm.DB {
	mConf := constant.Config.Mysql

	tools.CheckConfNil(mConf.Host, "Mysql-Host")
	tools.CheckConfNil(strconv.Itoa(mConf.Port), "Mysql-Port")
	tools.CheckConfNil(mConf.DbUser, "Mysql-DbUser")
	tools.CheckConfNil(mConf.DbPassword, "Mysql-DbPassword")
	tools.CheckConfNil(mConf.DbName, "Mysql-DbName")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mConf.DbUser,
			mConf.DbPassword,
			mConf.Host,
			strconv.Itoa(mConf.Port),
			mConf.DbName),
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		logrus.Errorln("数据库连接错误")
	} else {
		logrus.Infoln("数据库连接成功")
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	return db
}
