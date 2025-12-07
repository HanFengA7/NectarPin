package main

import (
	"log"
	"nectarpin/internal/api"
	"nectarpin/internal/config"
	"nectarpin/internal/database"
	"nectarpin/internal/models"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接
	if err := database.InitDatabase(cfg); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	defer func() {
		if err := database.CloseDatabase(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		}
	}()

	// 自动迁移数据库表结构（禁用外键）
	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化并启动服务器
	server := api.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
