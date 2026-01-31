package bootstrap

import (
	"fmt"
	"log"

	"NectarPin/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB
// @version 0.1.260126
// @desc 初始化数据库连接
// @param {Config} cfg 配置
// @returns {*gorm.DB} 数据库连接
// @throws {error} 连接失败
func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库
	if err := AutoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// AutoMigrate
// @version 0.1.260131
// @desc 自动迁移数据库表结构
// @param {*gorm.DB} db 数据库连接
// @returns {error} 迁移错误
func AutoMigrate(db *gorm.DB) error {
	log.Println("开始自动迁移数据库...")

	// 需要迁移的模型
	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Printf("数据库迁移失败: %v", err)
		return err
	}

	log.Println("数据库迁移完成")
	return nil
}
