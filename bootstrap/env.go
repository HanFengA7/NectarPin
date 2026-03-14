package bootstrap

import (
	"os"
	"path/filepath"
	"runtime"

	"nectarpin/internal/utils"

	"gopkg.in/yaml.v3"
)

// Config 应用程序配置结构体
// 包含服务器和数据库的所有配置项
type Config struct {
	Server   ServerConfig   // 服务器配置
	Database DatabaseConfig // 数据库配置
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port   int    `yaml:"port"`   // 服务监听端口
	Env    string `yaml:"env"`    // 运行环境 (development/production)
	Secret string `yaml:"secret"` // 应用密钥
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`     // 数据库主机地址
	Port     int    `yaml:"port"`     // 数据库端口
	User     string `yaml:"user"`     // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	Dbname   string `yaml:"dbname"`   // 数据库名称
	Schema   string `yaml:"schema"`   // 数据库 schema
}

// Env 环境配置封装
type Env struct {
	Env    string  // 当前运行环境
	Config *Config // 配置实例
}

// NewEnv 创建并加载环境配置
// 返回:
//   - *Env: 加载完成的环境配置实例
//
// 如果配置文件加载失败会直接终止程序
func NewEnv() *Env {
	config, err := loadConfig()
	if err != nil {
		utils.Logger.Fatalf("配置", "加载配置文件失败: %v", err)
	}

	return &Env{
		Env:    config.Server.Env,
		Config: config,
	}
}

// loadConfig 从配置文件加载配置
// 返回:
//   - *Config: 解析后的配置实例
//   - error: 加载过程中的错误信息
func loadConfig() (*Config, error) {
	configPath := getConfigPath()
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// getConfigPath 获取配置文件路径
// 返回:
//   - string: 配置文件的绝对路径
//
// 通过运行时信息定位项目根目录下的 config.yaml
func getConfigPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "config.yaml"
	}

	dir := filepath.Dir(filename)
	projectRoot := filepath.Join(dir, "..")

	return filepath.Join(projectRoot, "config.yaml")
}
