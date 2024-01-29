package PluginCore

import (
	"NectarPin/constant"
	"NectarPin/internal/PluginCore/PluginCorePB"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

type Service struct {
}

type PluginsList struct {
	PluginName         string
	PluginURL          string
	RouterName         string
	RouterNum          string
	RouterGroup        []*PluginCorePB.PluginRouterGroupRequest
	PluginInfo         PluginInfo
	PluginRouterStatus bool
}

type PluginInfo struct {
	PluginName    string
	PluginPort    string
	PluginVersion string
}

var PluginsListData []PluginsList

func extractParamsFromRoute(routePath string) ([]string, bool) {
	// 分割路径
	parts := strings.Split(routePath, "/")

	var params []string

	// 遍历路径中的部分，找到包含冒号的部分
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			// 去掉冒号
			param := strings.TrimLeft(part, ":")
			params = append(params, param)
		}
	}

	if len(params) > 0 {
		return params, true
	}

	return nil, false
}

func exponentialBackoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

func clearPluginStatus(pluginName string) {
	for i, v := range PluginsListData {
		if v.PluginName == pluginName {
			// 清理插件状态，确保重新注册时状态正确
			PluginsListData[i].PluginRouterStatus = false
			// 其他清理操作...
			break
		}
	}
}

// PluginRouteRegistered 插件路由注册 [PluginCore-grpcFunc][1]
func (s Service) PluginRouteRegistered(stream PluginCorePB.PluginService_PluginRouteRegisteredServer) error {
	//接收插件发来的信息
	request, _ := stream.Recv()

	var isInPLData1 = false
	for _, v := range PluginsListData {
		if v.PluginName == request.PluginInfo.PluginName {
			isInPLData1 = true
		}
	}

	if isInPLData1 == false {
		//将插件发来的信息存储到PluginsListData切片中
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
				PluginRouterStatus: true,
			})
		//路由注册及插件内部业务方法实现
		router := constant.Router
		for i, _ := range request.PluginRouterGroup {
			router.GET("/"+request.RouterName+"/"+request.PluginRouterGroup[i].RouterPath, func(c *gin.Context) {
				params, _ := extractParamsFromRoute(c.FullPath())
				// 动态获取params参数值
				var values []string
				for _, param := range params {
					value := c.Param(param)
					values = append(values, fmt.Sprintf("%s: %s", param, value))
				}
				// 将获取params参数值发給插件
				err := stream.Send(&PluginCorePB.PluginRouteRegisteredResponse{
					Code:       "200",
					Data2:      values,
					RouterPath: "/api/Plugins/" + request.RouterName,
					Message:    "路由发来的数据",
				})
				if err != nil {
					c.Abort()
					panic("[NectarPin-PluginCore]:Params参数值发給插件时出现异常！")
					return
				}
				// 接收插件处理后的数据

				//将接收插件处理后的数据发送给用户
				c.JSON(http.StatusOK, gin.H{"message": "Dynamic Route Added!",
					"params": values})
				time.Sleep(1 * time.Second)
			})
		}
	} else {
		for _, v := range PluginsListData {
			if v.PluginRouterStatus == true {
				fmt.Println("插件路由true")
			} else {
				fmt.Println("插件路由false")
				for i, v := range PluginsListData {
					if v.PluginName == request.PluginInfo.PluginName && v.PluginRouterStatus == false {
						fmt.Println(request.PluginInfo.PluginName + "\t<-插件被安装")
						PluginsListData[i].PluginRouterStatus = true
					}
				}
			}
		}
	}

	for i := 1; ; i++ {

		//发现插件卸载逻辑
		retryInterval := 10 * time.Second
		// 发起 gRPC 请求
		err1 := stream.Send(&PluginCorePB.PluginRouteRegisteredResponse{
			Code: "200",
			Data0: map[string]string{
				"TIME": time.Now().String(),
			},
			Message: "插件健康检测",
		})
		// gRPC 处理逻辑
		if err1 != nil {
			grpcStatus, ok := status.FromError(err1)
			if ok && grpcStatus.Code() == codes.Unavailable {
				// 处理连接不可用的情况，可以进行等待、重试等操作
				log.Println("gRPC connection unavailable:", grpcStatus.Err())
				time.Sleep(retryInterval)
				continue
			} else {
				// 其他错误处理
				log.Println("gRPC error:", err1)
			}
		}

		//	if err1 != nil {
		//		fmt.Println(err1)
		//		for _, v := range PluginsListData {
		//			if v.PluginName == request.PluginInfo.PluginName && v.PluginRouterStatus == true {
		//				fmt.Println(request.PluginInfo.PluginName + "\t<-插件被卸载")
		//				//PluginsListData = append(PluginsListData[:i], PluginsListData[i+1:]...)
		//				//PluginsListData[i].PluginRouterStatus = false
		//				// 在插件被卸载时清理状态
		//				clearPluginStatus(request.PluginInfo.PluginName)
		//			}
		//			break
		//		}
		//	}
		//	//backoffDuration := exponentialBackoff(i)
		//	//fmt.Printf("健康检查 %d: 下一次检测 %s\n", i+1, backoffDuration)
		//	time.Sleep(5 * time.Second)
		time.Sleep(15 * time.Second)
	}

}

//func (s Service) PluginInfo(ctx context.Context, request *PluginCorePB.PluginRequest) (*PluginCorePB.PluginResponse, error) {
//	return &PluginCorePB.PluginResponse{
//			PluginName: request.PluginName,
//			PluginURL:  request.PluginURL,
//		},
//		nil
//}
