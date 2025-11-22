package main

import (
	"log"
	"nectarpin/internal/api"
	"nectarpin/internal/config"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化并启动服务器
	server := api.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
