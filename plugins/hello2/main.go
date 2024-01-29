package main

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	ConfigExamplePlugin "NectarPin/plugins/hello/Config"
	InitExamplePlugin "NectarPin/plugins/hello/Init"
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	config := ConfigExamplePlugin.InitPluginConfig()
	configRouterNum, _ := strconv.Atoi(config.Router.RouterNum)

	var rGR []*PluginCorePB.PluginRouterGroupRequest
	// 使用 for 循环向 PluginRouterGroup 切片中添加元素
	for i := 1; i <= configRouterNum; i++ {
		// 获取结构体的字段通过反射获取结构体的反射值,使用直接字段访问方式获取字段的值
		routerGroupFieldValue := reflect.ValueOf(config.Router.RouterGroup).FieldByName(fmt.Sprintf("Routers%d", i))

		newRouterGroup := &PluginCorePB.PluginRouterGroupRequest{
			RouterPath:   fmt.Sprintf("%s", routerGroupFieldValue.FieldByName("RouterPath").String()),
			RouterMethod: fmt.Sprintf("%s", routerGroupFieldValue.FieldByName("RouterMethod").String()),
			RouterAuthIF: fmt.Sprintf("%s", routerGroupFieldValue.FieldByName("RouterAuthIF").String()),
		}

		rGR = append(rGR, newRouterGroup)
	}

	client := InitExamplePlugin.Client()
	stream, _ := client.PluginRouteRegistered(context.Background())
	_ = stream.Send(&PluginCorePB.PluginRouteRegisteredRequest{
		RouterName:        config.Router.RouterName,
		RouterNum:         config.Router.RouterNum,
		PluginRouterGroup: rGR,
		PluginInfo: &PluginCorePB.PluginInfoRequest{
			PluginName:    config.Plugin.PluginName,
			PluginPort:    config.Plugin.PluginPort,
			PluginVersion: config.Plugin.PluginVersion,
		},
	})

	// 开始异步接收数据
	go func() {
		// 从流中接收数据
		for {
			response, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}
			// 处理接收到的数据
			fmt.Printf("Received data: %v\n", response)
		}
	}()

	// 阻塞主线程，保持客户端运行
	select {}

}
