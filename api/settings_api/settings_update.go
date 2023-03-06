package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/response"
)

func (SettingsApi) SettingsUpdateView(c *gin.Context) {
	var cf config.SiteInfo
	err := c.ShouldBindJSON(&cf)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
	}
	global.Config.SiteInfo = cf //将前端传过来的值给到global.Config.SiteInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWith(c)
}
