package bootstrap

import (
	"fmt"
	"time"

	"nectarpin/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database 封装数据库连接实例
// 提供 GORM 数据库连接的管理和生命周期控制
type Database struct {
	DB *gorm.DB
}

// NewDatabase 创建并初始化数据库连接
// 参数:
//   - env: 环境配置实例，包含数据库连接参数
//
// 返回:
//   - *Database: 初始化完成的数据库实例
//
// 如果连接失败会直接终止程序
func NewDatabase(env *Env) *Database {
	db, err := connectToDatabase(env.Config.Database)
	if err != nil {
		utils.Logger.Fatalf("数据库", "连接数据库失败: %v", err)
	}

	return &Database{
		DB: db,
	}
}

// connectToDatabase 建立与 PostgreSQL 数据库的连接
// 参数:
//   - config: 数据库配置参数
//
// 返回:
//   - *gorm.DB: GORM 数据库实例
//   - error: 连接过程中的错误信息
//
// 该函数会:
//  1. 构建 DSN 连接字符串
//  2. 设置数据库 schema
//  3. 配置连接池参数
//  4. 验证数据库连接
func connectToDatabase(config DatabaseConfig) (*gorm.DB, error) {
	schema := config.Schema
	if schema == "" {
		schema = "public"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.User,
		config.Password,
		config.Dbname,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, utils.WrapError("打开数据库连接失败", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, utils.WrapError("获取底层 sql.DB 失败", err)
	}

	var currentSchema string
	err = sqlDB.QueryRow("SHOW search_path").Scan(&currentSchema)
	if err != nil {
		return nil, utils.WrapError("查询 search_path 失败", err)
	}
	utils.Logger.Infof("数据库", "默认 search_path: %s", currentSchema)

	_, err = sqlDB.Exec(fmt.Sprintf("SET search_path TO %s", schema))
	if err != nil {
		return nil, utils.WrapError("设置 search_path 失败", err)
	}

	err = sqlDB.QueryRow("SHOW search_path").Scan(&currentSchema)
	if err != nil {
		return nil, utils.WrapError("验证 search_path 失败", err)
	}
	utils.Logger.Infof("数据库", "设置后 search_path: %s", currentSchema)

	var schemaExists bool
	err = sqlDB.QueryRow("SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)", schema).Scan(&schemaExists)
	if err != nil {
		return nil, utils.WrapError("检查 schema 是否存在失败", err)
	}

	if !schemaExists {
		return nil, fmt.Errorf("schema '%s' 不存在，请先在数据库中创建", schema)
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)

	err = sqlDB.Ping()
	if err != nil {
		return nil, utils.WrapError("数据库 ping 失败", err)
	}

	utils.Logger.Successf("数据库", "数据库连接成功 (schema: %s)", schema)

	return db, nil
}

// Close 关闭数据库连接
// 返回:
//   - error: 关闭过程中的错误信息
//
// 应在应用程序关闭时调用以释放资源
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return utils.WrapError("获取底层 sql.DB 失败", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return utils.WrapError("关闭数据库连接失败", err)
	}

	utils.Logger.Info("数据库", "数据库连接已关闭")
	return nil
}
