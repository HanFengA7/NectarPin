package PluginCore

import (
	"math"
	"strings"
	"time"
)

func extractParamsFromRoute(routePath string) ([]string, bool) {
	// 分割路径
	parts := strings.Split(routePath, "/")

	var params []string

	// 遍历路径中的部分，找到包含冒号的部分
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			// 去掉冒号
			param := strings.TrimLeft(part, ":")
			params = append(params, param)
		}
	}

	if len(params) > 0 {
		return params, true
	}

	return nil, false
}

func exponentialBackoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

func clearPluginStatus(pluginName string) {
	for i, v := range PluginsListData {
		if v.PluginName == pluginName {
			// 清理插件状态，确保重新注册时状态正确
			PluginsListData[i].PluginRouterStatus = false
			// 其他清理操作...
			break
		}
	}
}
