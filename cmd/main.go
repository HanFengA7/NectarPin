// Package main 应用程序入口
// NectarPin - 一个现代化的 Web 应用后端服务
package main

import (
	"nectarpin/bootstrap"
)

// main 应用程序主入口函数
// 初始化并启动应用程序
func main() {
	app := bootstrap.App()
	app.Run()
}
