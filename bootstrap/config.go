package bootstrap

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port      int    `yaml:"port"`      // 端口
	Host      string `yaml:"host"`      // 主机
	TokenHash string `yaml:"tokenHash"` // Token哈希
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`     // 数据库主机
	Port     int    `yaml:"port"`     // 数据库端口
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	DBName   string `yaml:"dbName"`   // 数据库名
}

// LoadConfig
// @version 0.1.260126
// @desc 从指定路径读取并解析配置文件
// @param {string} path 配置文件路径
// @returns {*Config} 解析后的配置
// @throws {error} 读取或解析失败
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
