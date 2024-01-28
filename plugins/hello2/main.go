package main

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func initClient() PluginCorePB.PluginServiceClient {
	addr := ":3002"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	client := PluginCorePB.NewPluginServiceClient(conn)
	return client
}

func main() {
	//client := initClient()
	//time.Sleep(2000)
	//result, err := client.PluginInfo(context.Background(), &PluginCorePB.PluginRequest{
	//	PluginName: "Hello2插件",
	//	PluginURL:  ":3002",
	//})
	//result1, err := client.PluginRouteRegistered(context.Background(), &PluginCorePB.PluginRouterRequest{
	//	RouterPath: "hello2",
	//})
	//
	//fmt.Println(result, err)
	//fmt.Println(result1, err)
}
