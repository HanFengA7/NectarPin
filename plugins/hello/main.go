package main

import (
	"NectarPin/internal/PluginCore/PluginCorePB"
	ConfigExamplePlugin "NectarPin/plugins/hello/Config"
	InitExamplePlugin "NectarPin/plugins/hello/Init"
	"context"
	"fmt"
)

func main() {
	config := ConfigExamplePlugin.InitPluginConfig()
	fmt.Println(config)
	client := InitExamplePlugin.Client()
	stream, err := client.PluginRouteRegistered(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	err = stream.Send(&PluginCorePB.PluginRouteRegisteredRequest{
		RouterName:   "CosanoxjBot",
		RouterMethod: "GET",
		RouterAuthIF: false,
		PluginInfo: &PluginCorePB.PluginInfoRequest{
			PluginName:    "CosanoxjBot",
			PluginPort:    ":3002",
			PluginVersion: "v1.0.0",
		},
	})
	response, err := stream.Recv()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
