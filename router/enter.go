package router

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	engine := gin.Default()

	//路由分组
	apiRouterGroup := engine.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	//路由分层
	//系统配置api
	routerGroupApp.SettingsRouter()
	return engine
}
