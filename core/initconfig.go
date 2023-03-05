package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

// InitConfig 读取yaml配置文件
func InitConfig() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConfig, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConfig error: %s", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatal("config init Unmarshal: ", err)
	}
	log.Println("config yamlFile load init success.")
	global.Config = c
}
