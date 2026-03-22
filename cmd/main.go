// Package main 应用程序入口
// NectarPin - 一个现代化的 Web 应用后端服务
package main

import (
	"fmt"
	"os"

	"nectarpin/bootstrap"
)

// main 应用程序主入口函数
// 初始化并启动应用程序
//
// 命令行:
//   -init / --init  仅执行初始化向导写入 config.yaml，然后退出
//
// 若未带 -init 且项目根目录不存在 config.yaml，会自动进入向导再启动服务。
func main() {
	for _, a := range os.Args[1:] {
		if a == "-init" || a == "--init" {
			path := bootstrap.ConfigFilePath()
			if err := bootstrap.RunInitWizard(path); err != nil {
				fmt.Fprintf(os.Stderr, "初始化失败: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}
	}

	if !bootstrap.ConfigFileExists() {
		fmt.Println("未找到 config.yaml，启动初始化向导…")
		path := bootstrap.ConfigFilePath()
		if err := bootstrap.RunInitWizard(path); err != nil {
			fmt.Fprintf(os.Stderr, "初始化失败: %v\n", err)
			os.Exit(1)
		}
	}

	app := bootstrap.App()
	app.Run()
}
