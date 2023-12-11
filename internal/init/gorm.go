package init

import (
	"NectarPin/constant"
	"NectarPin/internal/models"
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
		constant.Log.Errorln("数据库连接错误")
		logrus.Errorln("数据库连接错误")
	} else {
		constant.Log.Info("数据库连接成功")
		logrus.Infoln("数据库连接成功")
	}

	//自动迁移
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Article{},
		&models.Comment{},
		&models.System{},
		&models.Link{},
	)
	if err != nil {
		return nil
	}
	/*
		add1 := models.User{
			Username:            "admin1",
			NickName:            "admin",
			Password:            "1235678",
			Email:               "1@1.com",
			AvatarUrl:           "https://img.com/o.png",
			Role:                1,
			LastLonginIPAddress: "10.1.0.1",
			LastLonginDate:      "2023-12-05",
		}
		add2 := models.Category{
			ParentID: 2,
			LevelNum: 1,
			Name:     "test2",
			Desc:     "test",
		}
		add3 := models.Article{
			User:          models.User{},
			Category:      models.Category{},
			UID:           1,
			CID:           2,
			Title:         "test2",
			Tags:          "test",
			Desc:          "test",
			Content:       "test",
			ImgUrl:        "test",
			AIFComment:    1,
			AIFHide:       0,
			AIFEncrypt:    0,
			AIFEncryptPwd: "",
		}
		db.Create(&add1)
		db.Create(&add2)
		db.Create(&add3)*/

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	return db
}
