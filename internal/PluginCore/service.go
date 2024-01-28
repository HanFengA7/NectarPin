package PluginCore

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"context"
)

type Service struct {
}

type PluginsList struct {
	PluginName   string
	PluginURL    string
	RouterAuthIF bool
}

var PluginsListData []PluginsList

func (s Service) PluginRouteRegistered(stream PluginCorePB.PluginService_PluginRouteRegisteredServer) error {

	request, _ := stream.Recv()
	PluginsListData = append(PluginsListData, PluginsList{PluginName: request.RouterName, PluginURL: "/api/Plugins/" + request.RouterName, RouterAuthIF: request.RouterAuthIF})

	//// 打印所有插件信息
	//for _, plugin := range PluginsListData {
	//	fmt.Printf("PluginName: %s, PluginURL: %s, RouterAuthIF:%s \n", plugin.PluginName, plugin.PluginURL, plugin.RouterAuthIF)
	//}

	err := stream.Send(&PluginCorePB.PluginRouteRegisteredResponse{Data: "200"})
	if err != nil {
		return err
	}
	return nil
}

func (s Service) PluginInfo(ctx context.Context, request *PluginCorePB.PluginRequest) (*PluginCorePB.PluginResponse, error) {
	return &PluginCorePB.PluginResponse{
			PluginName: request.PluginName,
			PluginURL:  request.PluginURL,
		},
		nil
}
