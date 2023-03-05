package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("警告！！！")
	global.Log.Error("错误！！！")
	global.Log.Infof("信息！！！")
	//连接数据库
	global.DB = core.InitGorm()

}
