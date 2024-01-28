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
	stream, err := client.PluginRouteRegistered(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	_ = stream.Send(&PluginCorePB.PluginRouteRegisteredRequest{
		RouterName:   "MengklBot",
		RouterAuthIF: false,
	})
	response, err := stream.Recv()
	fmt.Println(response, err)
}
