package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host      string
	Port      string
	Mode      string // debug, release, test
	JWTSecret string // JWT 密钥
	JWTExpire int    // JWT 过期时间（小时）
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string // PostgreSQL SSL模式: disable, require, verify-ca, verify-full
}

// Load 加载配置
// 优先级：YAML配置文件 > 默认值
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Host:      "0.0.0.0",
			Port:      "8080",
			Mode:      "debug",
			JWTSecret: "your-secret-key",
			JWTExpire: 24, // 默认24小时
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "",
			DBName:   "nectarpin",
			SSLMode:  "disable",
		},
	}

	// 尝试从YAML文件加载配置
	if err := loadFromYAML(cfg); err != nil {
		// YAML文件不存在或解析失败时，继续使用默认值
		// 这是正常的，所以不返回错误，但记录日志以便调试
		fmt.Printf("警告: 无法从YAML加载配置: %v，将使用默认值\n", err)
	} else {
		fmt.Printf("成功从YAML文件加载配置\n")
	}

	return cfg, nil
}

// findProjectRoot 向上查找项目根目录（包含go.mod的目录）
func findProjectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	current := wd
	for {
		// 检查当前目录是否有go.mod文件
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			return current
		}

		// 向上查找父目录
		parent := filepath.Dir(current)
		if parent == current {
			// 已经到达根目录
			break
		}
		current = parent
	}

	// 如果找不到，返回当前工作目录
	return wd
}

// loadFromYAML 从YAML文件加载配置
func loadFromYAML(cfg *Config) error {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	// 查找项目根目录
	projectRoot := findProjectRoot()

	// 查找配置文件，按优先级顺序
	configPaths := []string{
		filepath.Join(projectRoot, "configs", "config.yaml"),
		filepath.Join(projectRoot, "config.yaml"),
		filepath.Join(wd, "configs", "config.yaml"),
		filepath.Join(wd, "config.yaml"),
		"configs/config.yaml",
		"config.yaml",
		"./configs/config.yaml",
		"./config.yaml",
	}

	// 尝试从可执行文件所在目录查找
	if exePath, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exePath)
		configPaths = append([]string{
			filepath.Join(exeDir, "configs", "config.yaml"),
			filepath.Join(exeDir, "config.yaml"),
		}, configPaths...)
	}

	var configPath string
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			absPath, _ := filepath.Abs(path)
			configPath = absPath
			break
		}
	}

	if configPath == "" {
		return fmt.Errorf("配置文件不存在 (当前工作目录: %s, 项目根目录: %s)", wd, projectRoot)
	}

	fmt.Printf("找到配置文件: %s\n", configPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析YAML
	var yamlConfig struct {
		Server struct {
			Host      string `yaml:"host"`
			Port      string `yaml:"port"`
			Mode      string `yaml:"mode"`
			JWTSecret string `yaml:"jwt_secret"`
			JWTExpire int    `yaml:"jwt_expire"`
		} `yaml:"server"`
		Database struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
			SSLMode  string `yaml:"ssl_mode"`
		} `yaml:"database"`
	}

	if err := yaml.Unmarshal(data, &yamlConfig); err != nil {
		return fmt.Errorf("解析YAML配置失败: %w", err)
	}

	// 使用YAML配置（如果存在）
	if yamlConfig.Server.Host != "" {
		cfg.Server.Host = yamlConfig.Server.Host
	}
	if yamlConfig.Server.Port != "" {
		cfg.Server.Port = yamlConfig.Server.Port
	}
	if yamlConfig.Server.Mode != "" {
		cfg.Server.Mode = yamlConfig.Server.Mode
	}
	if yamlConfig.Server.JWTSecret != "" {
		cfg.Server.JWTSecret = yamlConfig.Server.JWTSecret
	}
	if yamlConfig.Server.JWTExpire > 0 {
		cfg.Server.JWTExpire = yamlConfig.Server.JWTExpire
	}

	if yamlConfig.Database.Host != "" {
		cfg.Database.Host = yamlConfig.Database.Host
	}
	if yamlConfig.Database.Port != "" {
		cfg.Database.Port = yamlConfig.Database.Port
	}
	if yamlConfig.Database.User != "" {
		cfg.Database.User = yamlConfig.Database.User
	}
	if yamlConfig.Database.Password != "" {
		cfg.Database.Password = yamlConfig.Database.Password
	}
	if yamlConfig.Database.DBName != "" {
		cfg.Database.DBName = yamlConfig.Database.DBName
	}
	if yamlConfig.Database.SSLMode != "" {
		cfg.Database.SSLMode = yamlConfig.Database.SSLMode
	}

	fmt.Printf("从YAML加载的数据库配置: host=%s, port=%s, user=%s, dbname=%s\n",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName)

	return nil
}
