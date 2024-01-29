package Init_ExamplePlugin

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
)

func Client() PluginCorePB.PluginServiceClient {
	addr := ":3002"
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(), grpc.WithBlock(), grpc.FailOnNonTempDialError(true),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                0,
				Timeout:             0,
				PermitWithoutStream: true,
			},
		),
	)
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	client := PluginCorePB.NewPluginServiceClient(conn)
	return client
}
