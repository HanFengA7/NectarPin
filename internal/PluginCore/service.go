package PluginCore

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	"fmt"
)

type Service struct {
}

type PluginsList struct {
	PluginName  string
	PluginURL   string
	RouterName  string
	RouterNum   string
	RouterGroup []*PluginCorePB.PluginRouterGroupRequest
	PluginInfo  PluginInfo
}

type PluginInfo struct {
	PluginName    string
	PluginPort    string
	PluginVersion string
}

var PluginsListData []PluginsList

func (s Service) PluginRouteRegistered(stream PluginCorePB.PluginService_PluginRouteRegisteredServer) error {

	//接收插件发来的信息
	request, _ := stream.Recv()
	//业务逻辑
	PluginsListData = append(PluginsListData,
		PluginsList{
			PluginName:  request.PluginInfo.PluginName,
			PluginURL:   "/api/Plugins/" + request.RouterName,
			RouterName:  request.RouterName,
			RouterNum:   request.RouterNum,
			RouterGroup: request.PluginRouterGroup,
			PluginInfo: PluginInfo{
				PluginName:    request.PluginInfo.PluginName,
				PluginPort:    request.PluginInfo.PluginPort,
				PluginVersion: request.PluginInfo.PluginVersion,
			},
		})
	fmt.Println(PluginsListData)
	//// 打印所有插件信息
	//for _, plugin := range PluginsListData {
	//	fmt.Printf("PluginName: %s, PluginURL: %s, RouterAuthIF:%s \n", plugin.PluginName, plugin.PluginURL, plugin.RouterAuthIF)
	//}

	//发送信息给插件
	err := stream.Send(&PluginCorePB.PluginRouteRegisteredResponse{
		Code: "200",
		Data: map[string]string{
			"key1": "ok",
			"key2": "yes",
		},
		RouterPath: "/api/Plugins/" + request.RouterName,
		Message:    "插件注册成功",
	})
	if err != nil {
		return err
	}
	return nil
}

//func (s Service) PluginInfo(ctx context.Context, request *PluginCorePB.PluginRequest) (*PluginCorePB.PluginResponse, error) {
//	return &PluginCorePB.PluginResponse{
//			PluginName: request.PluginName,
//			PluginURL:  request.PluginURL,
//		},
//		nil
//}
