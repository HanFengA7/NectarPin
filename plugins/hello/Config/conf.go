package Config_ExamplePlugin

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type PluginConfig struct {
	Plugin Plugin `yaml:"plugin"`
	Router Router `yaml:"router"`
}

type Plugin struct {
	PluginName    string `yaml:"pluginName"`
	PluginDesc    string `yaml:"pluginDesc"`
	PluginPort    int    `yaml:"pluginPort"`
	PluginVersion string `yaml:"PluginVersion"`
}

type Router struct {
	RouterName  string `yaml:"routerName"`
	RouterNum   int    `yaml:"routerNum"`
	RouterGroup struct {
		Routers1 Routers1 `yaml:"routers1"`
		Routers2 Routers2 `yaml:"routers2"`
		Routers3 Routers3 `yaml:"routers3"`
	} `yaml:"routerGroup"`
}

type Routers1 struct {
	RouterPath   string `yaml:"routerPath"`
	RouterMethod string `yaml:"routerMethod"`
	RouterAuthIF bool   `yaml:"routerAuthIF"`
}
type Routers2 struct {
	RouterPath   string `yaml:"routerPath"`
	RouterMethod string `yaml:"routerMethod"`
	RouterAuthIF bool   `yaml:"routerAuthIF"`
}
type Routers3 struct {
	RouterPath   string `yaml:"routerPath"`
	RouterMethod string `yaml:"routerMethod"`
	RouterAuthIF bool   `yaml:"routerAuthIF"`
}

func InitPluginConfig() *PluginConfig {
	config := &PluginConfig{}
	yamlConf, err := os.ReadFile("plugin.yaml")
	if err != nil {
		panic(fmt.Errorf("[NectarPin Error]: 配置文件路径错误"))
	}
	err = yaml.Unmarshal(yamlConf, config)

	if err != nil {
		logrus.Error(err)
		log.Fatalln("[NectarPin Error]: 配置文件Unmarshal错误")
	}
	return config
}
