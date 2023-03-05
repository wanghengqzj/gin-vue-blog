package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/router"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在：%s", addr)
	router.Run(addr)

}
