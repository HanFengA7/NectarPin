package PluginCore

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"context"
	"fmt"
)

type Service struct {
}

func (s Service) PluginRouteRegistered(ctx context.Context, request *PluginCorePB.PluginRouteRegisteredRequest) (
	*PluginCorePB.PluginRouteRegisteredResponse, error) {
	fmt.Println(request)

	//路由注册
	// [0] 判断路由是否要鉴权

	return &PluginCorePB.PluginRouteRegisteredResponse{
		Data: "200<--这是插件中心发来的",
	}, nil
}

func (s Service) PluginInfo(ctx context.Context, request *PluginCorePB.PluginRequest) (*PluginCorePB.PluginResponse, error) {
	return &PluginCorePB.PluginResponse{
			PluginName: request.PluginName,
			PluginURL:  request.PluginURL,
		},
		nil
}
