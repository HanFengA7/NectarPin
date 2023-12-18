package Init

import (
	"NectarPin/conf"
	"NectarPin/constant"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Conf() {
	const ConfigFile = "config.yaml"
	config := &conf.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("[NectarPin Error]: 配置文件路径错误"))
	}
	err = yaml.Unmarshal(yamlConf, config)

	if err != nil {
		logrus.Error(err)
		log.Fatalln("[NectarPin Error]: 配置文件Unmarshal错误")
	}
	constant.Config = config
}
