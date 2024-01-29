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
	stream, err := client.PluginRouteRegistered(context.Background())
	if err != nil {
		log.Fatalf("Error creating bidirectional stream: %v", err)
	}

	// 发送注册插件消息到服务端
	if err := stream.Send(&PluginCorePB.PluginRouteRegisteredRequest{
		RouterName:        config.Router.RouterName,
		RouterNum:         config.Router.RouterNum,
		PluginRouterGroup: rGR,
		PluginInfo: &PluginCorePB.PluginInfoRequest{
			PluginName:    config.Plugin.PluginName,
			PluginPort:    config.Plugin.PluginPort,
			PluginVersion: config.Plugin.PluginVersion,
		},
	}); err != nil {
		log.Fatalf("Error sending message to server: %v", err)
	}

	// 从流中接收数据
	for {
		// 发送心跳包消息到服务端
		//if err := stream.Send(&PluginCorePB.PluginRouteRegisteredRequest{
		//	RouterName: "这是心跳包[插件]",
		//}); err != nil {
		//	log.Fatalf("Error sending message to server: %v", err)
		//}
		//time.Sleep(5 * time.Second)

		// 接收服务端的响应
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving response from server: %v", err)
		}
		// 打印服务端的响应
		fmt.Printf("Server response: %v\n", response.Message)
	}

	//// 开始异步接收数据
	////go func() {
	//// 从流中接收数据
	//for {
	//	err1 := stream.Send(&PluginCorePB.PluginRouteRegisteredRequest{
	//		RouterName:        config.Router.RouterName,
	//		RouterNum:         config.Router.RouterNum,
	//		PluginRouterGroup: rGR,
	//		PluginInfo: &PluginCorePB.PluginInfoRequest{
	//			PluginName:    config.Plugin.PluginName,
	//			PluginPort:    config.Plugin.PluginPort,
	//			PluginVersion: config.Plugin.PluginVersion,
	//		},
	//	})
	//	fmt.Println(err1)
	//
	//	response, err := stream.Recv()
	//	if err != nil {
	//		log.Fatalf("Error receiving data: %v", err)
	//	}
	//	// 处理接收到的数据
	//	data := response
	//	fmt.Println(data)
	//}
	////}()
	////
	////// 阻塞主线程，保持客户端运行
	////select {}

}
