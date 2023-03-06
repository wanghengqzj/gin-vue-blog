package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConfig 读取yaml配置文件
func InitConfig() {
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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return nil
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return nil
	}
	global.Log.Info("配置文件修改成功！")
	return nil
}
