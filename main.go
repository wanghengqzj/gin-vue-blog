package main

import (
	"gvb_server/core"
	"gvb_server/flag"
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

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	engine := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在：%s", addr)
	err := engine.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
