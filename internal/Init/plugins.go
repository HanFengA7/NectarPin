package Init

import (
	"NectarPin/constant"
	"NectarPin/internal/PluginCore"
	"NectarPin/internal/PluginCore/PluginCorePB"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
	"net"
)

func Plugins() {

	//监听端口
	listen, err := net.Listen("tcp", ":3002")
	if err != nil {
		grpclog.Errorln(err)
	}
	//创建一个GRPC服务器实例
	gRPCServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     0,
		MaxConnectionAge:      0,
		MaxConnectionAgeGrace: 0,
		Time:                  0,
		Timeout:               0,
	}))
	server := PluginCore.Service{}
	// 将server结构体注册为gRPC服务。
	PluginCorePB.RegisterPluginServiceServer(gRPCServer, &server)
	constant.Log.Info("插件中心启动成功")
	logrus.Infoln("插件中心启动成功")
	// 开始处理客户端请求。
	err = gRPCServer.Serve(listen)
	if err != nil {
		grpclog.Errorln("服务失败:", err)
	}
}
