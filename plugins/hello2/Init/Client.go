package Init_ExamplePlugin

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func Client() PluginCorePB.PluginServiceClient {
	addr := ":3002"
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                5 * time.Minute, // 发送 ping 的频率
				PermitWithoutStream: true,            // 允许在没有活动流时保持连接
			},
		),
	)
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	client := PluginCorePB.NewPluginServiceClient(conn)
	return client
}
