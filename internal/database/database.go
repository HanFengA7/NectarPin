package database

import (
	"fmt"
	"log"
	"nectarpin/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// InitDatabase 初始化数据库连接
func InitDatabase(cfg *config.Config) error {
	// 打印配置信息以便调试（隐藏密码）
	log.Printf("数据库配置: host=%s, port=%s, user=%s, dbname=%s, sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.SSLMode)

	// 验证必要字段
	if cfg.Database.DBName == "" {
		return fmt.Errorf("数据库名称(DBName)不能为空")
	}
	if cfg.Database.Host == "" {
		return fmt.Errorf("数据库主机(Host)不能为空")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	log.Printf("数据库连接DSN: host=%s, user=%s, dbname=%s, port=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.DBName, cfg.Database.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 使用复数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})

	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 测试数据库连接
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}

	log.Printf("数据库连接成功: %s@%s:%s/%s", cfg.Database.User, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	return nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("数据库未初始化")
	}

	log.Println("开始自动迁移数据库表结构...")
	err := DB.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}

	log.Println("数据库表结构迁移完成")
	return nil
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
