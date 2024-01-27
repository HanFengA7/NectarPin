package main

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func initClient() PluginCorePB.PluginServiceClient {
	addr := ":3002"
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:    10 * time.Second, // 发送周期，表示每隔多长时间发送一个 Ping
				Timeout: 5 * time.Second,  // 超时时间，表示多长时间没有收到 Ping 就认为连接断开
			},
		),
	)
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	client := PluginCorePB.NewPluginServiceClient(conn)
	return client
}

func main() {
	client := initClient()
	//result, err := client.PluginInfo(context.Background(), &PluginCorePB.PluginRequest{
	//	PluginName: "Hello插件",
	//	PluginURL:  ":3002",
	//})
	for i := 0; i < 3; i++ {
		result1, err := client.PluginRouteRegistered(context.Background(), &PluginCorePB.PluginRouteRegisteredRequest{
			RouterName:   "hello",
			RouterAuthIF: false,
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result1)
		time.Sleep(5000)
	}

}
